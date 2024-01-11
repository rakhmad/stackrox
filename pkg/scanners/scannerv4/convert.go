package scannerv4

import (
	"fmt"
	"strings"

	gogotypes "github.com/gogo/protobuf/types"
	v4 "github.com/stackrox/rox/generated/internalapi/scanner/v4"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/clair"
	"github.com/stackrox/rox/pkg/cvss/cvssv2"
	"github.com/stackrox/rox/pkg/cvss/cvssv3"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/utils"
)

func imageScan(metadata *storage.ImageMetadata, report *v4.VulnerabilityReport) *storage.ImageScan {
	scan := &storage.ImageScan{
		// TODO(ROX-21362): Get ScannerVersion from ScannerV4 matcher API
		// ScannerVersion: ,
		ScanTime:        gogotypes.TimestampNow(),
		OperatingSystem: os(report),
		Components:      components(metadata, report),
	}

	if scan.GetOperatingSystem() == "unknown" {
		scan.Notes = append(scan.Notes, storage.ImageScan_OS_UNAVAILABLE)
	}

	return scan
}

func components(metadata *storage.ImageMetadata, report *v4.VulnerabilityReport) []*storage.EmbeddedImageScanComponent {
	layerSHAToIndex := clair.BuildSHAToIndexMap(metadata)

	components := make([]*storage.EmbeddedImageScanComponent, 0, len(report.GetPackageVulnerabilities()))
	for _, pkg := range report.GetContents().GetPackages() {
		id := pkg.GetId()
		vulnIDs := report.GetPackageVulnerabilities()[id].GetValues()
		component := &storage.EmbeddedImageScanComponent{
			Name:          pkg.GetName(),
			Version:       pkg.GetVersion(),
			Vulns:         vulnerabilities(report.GetVulnerabilities(), vulnIDs),
			Location:      pkg.GetPackageDb(),
			HasLayerIndex: layerIndex(layerSHAToIndex, report, id),
		}

		components = append(components, component)
	}

	return components
}

func layerIndex(layerSHAToIndex map[string]int32, report *v4.VulnerabilityReport, pkgID string) *storage.EmbeddedImageScanComponent_LayerIndex {
	// TODO(ROX-21377): Confirm with Clair team how handle multiple environments
	envList := report.GetContents().GetEnvironments()[pkgID]
	if len(envList.GetEnvironments()) > 0 {
		env := envList.GetEnvironments()[0]

		if val, ok := layerSHAToIndex[env.GetIntroducedIn()]; ok {
			return &storage.EmbeddedImageScanComponent_LayerIndex{
				LayerIndex: val,
			}
		}
	}

	return nil
}

func vulnerabilities(vulnerabilities map[string]*v4.VulnerabilityReport_Vulnerability, ids []string) []*storage.EmbeddedVulnerability {
	if len(vulnerabilities) == 0 {
		return nil
	}

	vulns := make([]*storage.EmbeddedVulnerability, 0, len(ids))
	uniqueVulns := set.NewStringSet()
	for _, id := range ids {
		ccVuln, ok := vulnerabilities[id]
		if !ok {
			log.Debugf("Bad Input: Vuln %q from PackageVulnerabilities not found in Vulnerabilities, skipping", id)
			continue
		}

		if !uniqueVulns.Add(ccVuln.Name) {
			// Already added this vulnerability, so ignore it.
			continue
		}

		// TODO(ROX-20355): Populate last modified once the API is available.
		vuln := &storage.EmbeddedVulnerability{
			Cve:         ccVuln.GetName(),
			Summary:     ccVuln.GetDescription(),
			Link:        link(ccVuln.GetLink()),
			PublishedOn: ccVuln.GetIssued(),
			// LastModified: ,
			VulnerabilityType: storage.EmbeddedVulnerability_IMAGE_VULNERABILITY,
			Severity:          normalizedSeverity(ccVuln.GetNormalizedSeverity()),
		}
		if err := setCvss(vuln, ccVuln.GetCvss()); err != nil {
			utils.Should(err)
		}
		if ccVuln.GetFixedInVersion() != "" {
			vuln.SetFixedBy = &storage.EmbeddedVulnerability_FixedBy{
				FixedBy: ccVuln.GetFixedInVersion(),
			}
		}

		vulns = append(vulns, vuln)
	}

	return vulns
}

func setCvss(vuln *storage.EmbeddedVulnerability, cvss *v4.VulnerabilityReport_Vulnerability_CVSS) error {
	if cvss == nil {
		return nil
	}
	errList := errorhelpers.NewErrorList("failed to parse vector")
	if v2 := cvss.GetV2(); v2 != nil {
		if c, err := cvssv2.ParseCVSSV2(v2.GetVector()); err == nil {
			err = cvssv2.CalculateScores(c)
			if err != nil {
				errList.AddError(fmt.Errorf("calculating CVSS v2 scores: %w", err))
			}
			// Use the report's score if it exists.
			if v2.GetBaseScore() != 0.0 {
				c.Score = v2.GetBaseScore()
			}
			c.Severity = cvssv2.Severity(v2.GetBaseScore())
			vuln.CvssV2 = c
			// This sets the top level score for use in policies. It will be overwritten if
			// v3 exists.
			vuln.ScoreVersion = storage.EmbeddedVulnerability_V2
			vuln.Cvss = v2.GetBaseScore()
		} else {
			errList.AddError(fmt.Errorf("v2: %w", err))
		}
	}
	if v3 := cvss.GetV3(); v3 != nil {
		if c, err := cvssv3.ParseCVSSV3(v3.GetVector()); err == nil {
			err = cvssv3.CalculateScores(c)
			if err != nil {
				errList.AddError(fmt.Errorf("calculating CVSS v3 scores: %w", err))
			}
			// Use the report's score if it exists.
			if v3.GetBaseScore() != 0.0 {
				c.Score = v3.GetBaseScore()
			}
			c.Severity = cvssv3.Severity(v3.GetBaseScore())
			vuln.CvssV3 = c
			// Overwrite V2 if set.
			vuln.ScoreVersion = storage.EmbeddedVulnerability_V3
			vuln.Cvss = v3.GetBaseScore()
		} else {
			errList.AddError(fmt.Errorf("v3: %w", err))
		}
	}
	return errList.ToError()
}

// link returns the first link from space separated list of links (which is how ClairCore provides links).
// The ACS UI will fail to show a vulnerability's link if it is an invalid URL.
func link(links string) string {
	link, _, _ := strings.Cut(links, " ")
	return link
}

func normalizedSeverity(severity v4.VulnerabilityReport_Vulnerability_Severity) storage.VulnerabilitySeverity {
	switch severity {
	case v4.VulnerabilityReport_Vulnerability_SEVERITY_LOW:
		return storage.VulnerabilitySeverity_LOW_VULNERABILITY_SEVERITY
	case v4.VulnerabilityReport_Vulnerability_SEVERITY_MODERATE:
		return storage.VulnerabilitySeverity_MODERATE_VULNERABILITY_SEVERITY
	case v4.VulnerabilityReport_Vulnerability_SEVERITY_IMPORTANT:
		return storage.VulnerabilitySeverity_IMPORTANT_VULNERABILITY_SEVERITY
	case v4.VulnerabilityReport_Vulnerability_SEVERITY_CRITICAL:
		return storage.VulnerabilitySeverity_CRITICAL_VULNERABILITY_SEVERITY
	default:
		return storage.VulnerabilitySeverity_UNKNOWN_VULNERABILITY_SEVERITY
	}
}

// os retrieves the OS name:version for the image represented by the given
// vulnerability report.
// If there are zero known distributions for the image or if there are multiple distributions,
// return "unknown", as StackRox only supports a single base-OS at this time.
func os(report *v4.VulnerabilityReport) string {
	if len(report.GetContents().GetDistributions()) == 1 {
		for _, dist := range report.GetContents().GetDistributions() {
			return dist.Did + ":" + dist.VersionId
		}
	}

	return "unknown"
}

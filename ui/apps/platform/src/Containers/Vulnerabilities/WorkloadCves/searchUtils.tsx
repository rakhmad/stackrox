import qs from 'qs';
import { vulnerabilitiesWorkloadCvesPath } from 'routePaths';
import { SearchFilter } from 'types/search';
import { getQueryString } from 'utils/queryStringUtils';

export type CveStatusTab = 'Observed' | 'Deferred' | 'False Positive';
export type EntityTab = 'CVE' | 'Image' | 'Deployment';

export type WorkloadCvesSearch = {
    cveStatusTab: CveStatusTab;
    entityTab?: EntityTab;
    s?: SearchFilter;
};

function isValidCveStatusTab(cveStatusTab: unknown): cveStatusTab is CveStatusTab {
    return (
        cveStatusTab === 'Observed' ||
        cveStatusTab === 'Deferred' ||
        cveStatusTab === 'False Positive'
    );
}

export function parseWorkloadCvesOverviewSearchString(search: string): WorkloadCvesSearch {
    const { cveStatusTab } = qs.parse(search, { ignoreQueryPrefix: true });

    return {
        cveStatusTab: isValidCveStatusTab(cveStatusTab) ? cveStatusTab : 'Observed',
    };
}

export function getOverviewCvesPath(workloadCvesSearch: WorkloadCvesSearch) {
    return `${vulnerabilitiesWorkloadCvesPath}${getQueryString(workloadCvesSearch)}`;
}

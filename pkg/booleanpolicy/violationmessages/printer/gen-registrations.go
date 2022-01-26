// Code generated by violationprinter codegen. DO NOT EDIT.

package printer

const (
	AddCapabilityKey                = "addCapability"
	AppArmorProfileKey              = "appArmorProfile"
	AutomountServiceAccountTokenKey = "automountServiceAccountToken"
	ComponentKey                    = "component"
	ContainerNameKey                = "containerName"
	CveKey                          = "cve"
	DisallowedAnnotationKey         = "disallowedAnnotation"
	DisallowedImageLabelKey         = "disallowedImageLabel"
	DropCapabilityKey               = "dropCapability"
	EnvKey                          = "env"
	HostIPCKey                      = "hostIPC"
	HostNetworkKey                  = "hostNetwork"
	HostPIDKey                      = "hostPID"
	ImageAgeKey                     = "imageAge"
	ImageDetailsKey                 = "imageDetails"
	ImageOSKey                      = "imageOS"
	ImageScanKey                    = "imageScan"
	ImageScanAgeKey                 = "imageScanAge"
	ImageUserKey                    = "imageUser"
	LineKey                         = "line"
	NamespaceKey                    = "namespace"
	NodePortKey                     = "nodePort"
	PortKey                         = "port"
	PortExposureKey                 = "portExposure"
	PrivilegedKey                   = "privileged"
	ProcessBaselineKey              = "processBaseline"
	RbacKey                         = "rbac"
	ReadOnlyRootFSKey               = "readOnlyRootFS"
	ReplicasKey                     = "replicas"
	RequiredAnnotationKey           = "requiredAnnotation"
	RequiredImageLabelKey           = "requiredImageLabel"
	RequiredLabelKey                = "requiredLabel"
	ResourceKey                     = "resource"
	RuntimeClassKey                 = "runtimeClass"
	SeccompProfileTypeKey           = "seccompProfileType"
	ServiceAccountKey               = "serviceAccount"
	VolumeKey                       = "volume"
)

func init() {
	registerFunc(AddCapabilityKey, addCapabilityPrinter)
	registerFunc(AppArmorProfileKey, appArmorProfilePrinter)
	registerFunc(AutomountServiceAccountTokenKey, automountServiceAccountTokenPrinter)
	registerFunc(ComponentKey, componentPrinter)
	registerFunc(ContainerNameKey, containerNamePrinter)
	registerFunc(CveKey, cvePrinter)
	registerFunc(DisallowedAnnotationKey, disallowedAnnotationPrinter)
	registerFunc(DisallowedImageLabelKey, disallowedImageLabelPrinter)
	registerFunc(DropCapabilityKey, dropCapabilityPrinter)
	registerFunc(EnvKey, envPrinter)
	registerFunc(HostIPCKey, hostIPCPrinter)
	registerFunc(HostNetworkKey, hostNetworkPrinter)
	registerFunc(HostPIDKey, hostPIDPrinter)
	registerFunc(ImageAgeKey, imageAgePrinter)
	registerFunc(ImageDetailsKey, imageDetailsPrinter)
	registerFunc(ImageOSKey, imageOSPrinter)
	registerFunc(ImageScanKey, imageScanPrinter)
	registerFunc(ImageScanAgeKey, imageScanAgePrinter)
	registerFunc(ImageUserKey, imageUserPrinter)
	registerFunc(LineKey, linePrinter)
	registerFunc(NamespaceKey, namespacePrinter)
	registerFunc(NodePortKey, nodePortPrinter)
	registerFunc(PortKey, portPrinter)
	registerFunc(PortExposureKey, portExposurePrinter)
	registerFunc(PrivilegedKey, privilegedPrinter)
	registerFunc(ProcessBaselineKey, processBaselinePrinter)
	registerFunc(RbacKey, rbacPrinter)
	registerFunc(ReadOnlyRootFSKey, readOnlyRootFSPrinter)
	registerFunc(ReplicasKey, replicasPrinter)
	registerFunc(RequiredAnnotationKey, requiredAnnotationPrinter)
	registerFunc(RequiredImageLabelKey, requiredImageLabelPrinter)
	registerFunc(RequiredLabelKey, requiredLabelPrinter)
	registerFunc(ResourceKey, resourcePrinter)
	registerFunc(RuntimeClassKey, runtimeClassPrinter)
	registerFunc(SeccompProfileTypeKey, seccompProfileTypePrinter)
	registerFunc(ServiceAccountKey, serviceAccountPrinter)
	registerFunc(VolumeKey, volumePrinter)
}

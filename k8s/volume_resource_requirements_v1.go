// Code generated; DO NOT EDIT.

package k8s

// VolumeResourceRequirementsV1 VolumeResourceRequirements describes the storage resource requirements for a volume.
type VolumeResourceRequirementsV1 struct {
	// Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits *map[string]string `json:"limits,omitempty"`
	// Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests *map[string]string `json:"requests,omitempty"`
}

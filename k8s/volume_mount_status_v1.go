// Code generated; DO NOT EDIT.

package k8s

// VolumeMountStatusV1 VolumeMountStatus shows status of volume mounts.
type VolumeMountStatusV1 struct {
	// MountPath corresponds to the original VolumeMount.
	MountPath string `json:"mountPath"`
	// Name corresponds to the name of the original VolumeMount.
	Name string `json:"name"`
	// ReadOnly corresponds to the original VolumeMount.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// RecursiveReadOnly must be set to Disabled, Enabled, or unspecified (for non-readonly mounts). An IfPossible value in the original VolumeMount must be translated to Disabled or Enabled, depending on the mount result.
	RecursiveReadOnly *string `json:"recursiveReadOnly,omitempty"`
}

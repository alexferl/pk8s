// Code generated by pk8s; DO NOT EDIT.

package k8s

// FlexPersistentVolumeSourceV1 FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.
type FlexPersistentVolumeSourceV1 struct {
	// driver is the name of the driver to use for this volume.
	Driver string `json:"driver"`
	// fsType is the Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	FSType *string `json:"fsType,omitempty"`
	// options is Optional: this field holds extra command options if any.
	Options *map[string]string `json:"options,omitempty"`
	// readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	SecretRef *SecretReferenceV1 `json:"secretRef,omitempty"`
}

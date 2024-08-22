// Code generated; DO NOT EDIT.

package k8s

// FlexVolumeSourceV1 FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
type FlexVolumeSourceV1 struct {
	// driver is the name of the driver to use for this volume.
	Driver string `json:"driver"`
	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	FsType *string `json:"fsType,omitempty"`
	// options is Optional: this field holds extra command options if any.
	Options *map[string]string `json:"options,omitempty"`
	// readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	SecretRef *LocalObjectReferenceV1 `json:"secretRef,omitempty"`
}

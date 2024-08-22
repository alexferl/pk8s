// Code generated; DO NOT EDIT.

package k8s

// PortworxVolumeSourceV1 PortworxVolumeSource represents a Portworx volume resource.
type PortworxVolumeSourceV1 struct {
	// fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// volumeID uniquely identifies a Portworx volume
	VolumeID string `json:"volumeID"`
}

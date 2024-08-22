// Code generated; DO NOT EDIT.

package k8s

// CSIVolumeSourceV1 Represents a source location of a volume to mount, managed by an external CSI driver
type CSIVolumeSourceV1 struct {
	// driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.
	Driver string `json:"driver"`
	// fsType to mount. Ex. "ext4", "xfs", "ntfs". If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.
	FsType *string `json:"fsType,omitempty"`
	// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	NodePublishSecretRef *LocalObjectReferenceV1 `json:"nodePublishSecretRef,omitempty"`
	// readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).
	ReadOnly *bool `json:"readOnly,omitempty"`
	// volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.
	VolumeAttributes *map[string]string `json:"volumeAttributes,omitempty"`
}

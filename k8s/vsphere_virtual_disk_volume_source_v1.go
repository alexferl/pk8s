// Code generated; DO NOT EDIT.

package k8s

// VsphereVirtualDiskVolumeSourceV1 Represents a vSphere volume resource.
type VsphereVirtualDiskVolumeSourceV1 struct {
	// fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	StoragePolicyID *string `json:"storagePolicyID,omitempty"`
	// storagePolicyName is the storage Policy Based Management (SPBM) profile name.
	StoragePolicyName *string `json:"storagePolicyName,omitempty"`
	// volumePath is the path that identifies vSphere volume vmdk
	VolumePath string `json:"volumePath"`
}

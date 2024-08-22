// Code generated; DO NOT EDIT.

package k8s

// PhotonPersistentDiskVolumeSourceV1 Represents a Photon Controller persistent disk resource.
type PhotonPersistentDiskVolumeSourceV1 struct {
	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// pdID is the ID that identifies Photon Controller persistent disk
	PdID string `json:"pdID"`
}

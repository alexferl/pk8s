// Code generated; DO NOT EDIT.

package k8s

// LocalVolumeSourceV1 Local represents directly-attached storage with node affinity (Beta feature)
type LocalVolumeSourceV1 struct {
	// fsType is the filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a filesystem if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// path of the full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).
	Path string `json:"path"`
}

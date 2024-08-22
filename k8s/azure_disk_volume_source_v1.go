// Code generated; DO NOT EDIT.

package k8s

// AzureDiskVolumeSourceV1 AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type AzureDiskVolumeSourceV1 struct {
	// cachingMode is the Host Caching mode: None, Read Only, Read Write.
	CachingMode *string `json:"cachingMode,omitempty"`
	// diskName is the Name of the data disk in the blob storage
	DiskName string `json:"diskName"`
	// diskURI is the URI of data disk in the blob storage
	DiskURI string `json:"diskURI"`
	// fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
	Kind *string `json:"kind,omitempty"`
	// readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

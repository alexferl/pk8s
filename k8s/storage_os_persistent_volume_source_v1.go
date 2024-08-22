// Code generated; DO NOT EDIT.

package k8s

// StorageOSPersistentVolumeSourceV1 Represents a StorageOS persistent volume resource.
type StorageOSPersistentVolumeSourceV1 struct {
	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// ObjectReference contains enough information to let you inspect or modify the referred object.
	SecretRef *ObjectReferenceV1 `json:"secretRef,omitempty"`
	// volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.
	VolumeName *string `json:"volumeName,omitempty"`
	// volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
	VolumeNamespace *string `json:"volumeNamespace,omitempty"`
}

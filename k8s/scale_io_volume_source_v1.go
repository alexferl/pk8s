// Code generated; DO NOT EDIT.

package k8s

// ScaleIOVolumeSourceV1 ScaleIOVolumeSource represents a persistent ScaleIO volume
type ScaleIOVolumeSourceV1 struct {
	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs".
	FsType *string `json:"fsType,omitempty"`
	// gateway is the host address of the ScaleIO API Gateway.
	Gateway string `json:"gateway"`
	// protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.
	ProtectionDomain *string `json:"protectionDomain,omitempty"`
	// readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	SecretRef LocalObjectReferenceV1 `json:"secretRef"`
	// sslEnabled Flag enable/disable SSL communication with Gateway, default false
	SslEnabled *bool `json:"sslEnabled,omitempty"`
	// storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
	StorageMode *string `json:"storageMode,omitempty"`
	// storagePool is the ScaleIO Storage Pool associated with the protection domain.
	StoragePool *string `json:"storagePool,omitempty"`
	// system is the name of the storage system as configured in ScaleIO.
	System string `json:"system"`
	// volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.
	VolumeName *string `json:"volumeName,omitempty"`
}

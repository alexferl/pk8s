// Code generated; DO NOT EDIT.

package k8s

// CSIPersistentVolumeSourceV1 Represents storage that is managed by an external CSI volume driver (Beta feature)
type CSIPersistentVolumeSourceV1 struct {
	// SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	ControllerExpandSecretRef *SecretReferenceV1 `json:"controllerExpandSecretRef,omitempty"`
	// SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	ControllerPublishSecretRef *SecretReferenceV1 `json:"controllerPublishSecretRef,omitempty"`
	// driver is the name of the driver to use for this volume. Required.
	Driver string `json:"driver"`
	// fsType to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	FsType *string `json:"fsType,omitempty"`
	// SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	NodeExpandSecretRef *SecretReferenceV1 `json:"nodeExpandSecretRef,omitempty"`
	// SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	NodePublishSecretRef *SecretReferenceV1 `json:"nodePublishSecretRef,omitempty"`
	// SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
	NodeStageSecretRef *SecretReferenceV1 `json:"nodeStageSecretRef,omitempty"`
	// readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
	ReadOnly *bool `json:"readOnly,omitempty"`
	// volumeAttributes of the volume to publish.
	VolumeAttributes *map[string]string `json:"volumeAttributes,omitempty"`
	// volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.
	VolumeHandle string `json:"volumeHandle"`
}

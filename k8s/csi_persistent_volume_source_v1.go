// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// CSIPersistentVolumeSourceV1 Represents storage that is managed by an external CSI volume driver (Beta feature)
type CSIPersistentVolumeSourceV1 struct {
	ControllerExpandSecretRef  *SecretReferenceV1 `json:"controllerExpandSecretRef,omitempty"`
	ControllerPublishSecretRef *SecretReferenceV1 `json:"controllerPublishSecretRef,omitempty"`
	// driver is the name of the driver to use for this volume. Required.
	Driver string `json:"driver"`
	// fsType to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\".
	FsType               *string            `json:"fsType,omitempty"`
	NodeExpandSecretRef  *SecretReferenceV1 `json:"nodeExpandSecretRef,omitempty"`
	NodePublishSecretRef *SecretReferenceV1 `json:"nodePublishSecretRef,omitempty"`
	NodeStageSecretRef   *SecretReferenceV1 `json:"nodeStageSecretRef,omitempty"`
	// readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
	ReadOnly *bool `json:"readOnly,omitempty"`
	// volumeAttributes of the volume to publish.
	VolumeAttributes *map[string]string `json:"volumeAttributes,omitempty"`
	// volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.
	VolumeHandle string `json:"volumeHandle"`
}

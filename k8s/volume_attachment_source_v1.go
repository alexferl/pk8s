// Code generated; DO NOT EDIT.

package k8s

// VolumeAttachmentSourceV1 VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
type VolumeAttachmentSourceV1 struct {
	// PersistentVolumeSpec is the specification of a persistent volume.
	InlineVolumeSpec *PersistentVolumeSpecV1 `json:"inlineVolumeSpec,omitempty"`
	// persistentVolumeName represents the name of the persistent volume to attach.
	PersistentVolumeName *string `json:"persistentVolumeName,omitempty"`
}

// Code generated by pk8s; DO NOT EDIT.

package k8s

// VolumeAttachmentSpecV1 VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpecV1 struct {
	// attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
	Attacher string `json:"attacher"`
	// nodeName represents the node that the volume should be attached to.
	NodeName string `json:"nodeName"`
	// VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
	Source VolumeAttachmentSourceV1 `json:"source"`
}

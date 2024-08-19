// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// VolumeAttachmentStatusV1 VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatusV1 struct {
	AttachError *VolumeErrorV1 `json:"attachError,omitempty"`
	// attached indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Attached bool `json:"attached"`
	// attachmentMetadata is populated with any information returned by the attach operation, upon successful attach, that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachmentMetadata *map[string]string `json:"attachmentMetadata,omitempty"`
	DetachError        *VolumeErrorV1     `json:"detachError,omitempty"`
}

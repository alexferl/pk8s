// Code generated; DO NOT EDIT.

package k8s

// PersistentVolumeClaimTemplateV1 PersistentVolumeClaimTemplate is used to produce PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
type PersistentVolumeClaimTemplateV1 struct {
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
	Spec PersistentVolumeClaimSpecV1 `json:"spec"`
}

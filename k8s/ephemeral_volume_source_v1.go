// Code generated; DO NOT EDIT.

package k8s

// EphemeralVolumeSourceV1 Represents an ephemeral volume that is handled by a normal storage driver.
type EphemeralVolumeSourceV1 struct {
	// PersistentVolumeClaimTemplate is used to produce PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
	VolumeClaimTemplate *PersistentVolumeClaimTemplateV1 `json:"volumeClaimTemplate,omitempty"`
}

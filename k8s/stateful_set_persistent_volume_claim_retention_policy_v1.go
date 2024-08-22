// Code generated; DO NOT EDIT.

package k8s

// StatefulSetPersistentVolumeClaimRetentionPolicyV1 StatefulSetPersistentVolumeClaimRetentionPolicy describes the policy used for PVCs created from the StatefulSet VolumeClaimTemplates.
type StatefulSetPersistentVolumeClaimRetentionPolicyV1 struct {
	// WhenDeleted specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is deleted. The default policy of `Retain` causes PVCs to not be affected by StatefulSet deletion. The `Delete` policy causes those PVCs to be deleted.
	WhenDeleted *string `json:"whenDeleted,omitempty"`
	// WhenScaled specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is scaled down. The default policy of `Retain` causes PVCs to not be affected by a scaledown. The `Delete` policy causes the associated PVCs for any excess pods above the replica count to be deleted.
	WhenScaled *string `json:"whenScaled,omitempty"`
}

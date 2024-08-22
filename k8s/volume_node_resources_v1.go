// Code generated; DO NOT EDIT.

package k8s

// VolumeNodeResourcesV1 VolumeNodeResources is a set of resource limits for scheduling of volumes.
type VolumeNodeResourcesV1 struct {
	// count indicates the maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is not specified, then the supported number of volumes on this node is unbounded.
	Count *int32 `json:"count,omitempty"`
}

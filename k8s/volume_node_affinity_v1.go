// Code generated; DO NOT EDIT.

package k8s

// VolumeNodeAffinityV1 VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
type VolumeNodeAffinityV1 struct {
	// A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
	Required *NodeSelectorV1 `json:"required,omitempty"`
}

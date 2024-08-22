// Code generated; DO NOT EDIT.

package k8s

// NodeSelectorTermV1 A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
type NodeSelectorTermV1 struct {
	// A list of node selector requirements by node's labels.
	MatchExpressions []NodeSelectorRequirementV1 `json:"matchExpressions,omitempty"`
	// A list of node selector requirements by node's fields.
	MatchFields []NodeSelectorRequirementV1 `json:"matchFields,omitempty"`
}

// Code generated by pk8s; DO NOT EDIT.

package k8s

// NodeAffinityV1 Node affinity is a group of node affinity scheduling rules.
type NodeAffinityV1 struct {
	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTermV1 `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
	// A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
	RequiredDuringSchedulingIgnoredDuringExecution *NodeSelectorV1 `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

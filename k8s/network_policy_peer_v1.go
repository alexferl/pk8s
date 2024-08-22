// Code generated; DO NOT EDIT.

package k8s

// NetworkPolicyPeerV1 NetworkPolicyPeer describes a peer to allow traffic to/from. Only certain combinations of fields are allowed
type NetworkPolicyPeerV1 struct {
	// IPBlock describes a particular CIDR (Ex. "192.168.1.0/24","2001:db8::/64") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
	IPBlock *IPBlockV1 `json:"ipBlock,omitempty"`
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	NamespaceSelector *LabelSelectorV1 `json:"namespaceSelector,omitempty"`
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	PodSelector *LabelSelectorV1 `json:"podSelector,omitempty"`
}

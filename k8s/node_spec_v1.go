// Code generated; DO NOT EDIT.

package k8s

// NodeSpecV1 NodeSpec describes the attributes that a node is created with.
type NodeSpecV1 struct {
	// NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil. This API is deprecated since 1.22
	ConfigSource *NodeConfigSourceV1 `json:"configSource,omitempty"`
	// Deprecated. Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966
	ExternalID *string `json:"externalID,omitempty"`
	// PodCIDR represents the pod IP range assigned to the node.
	PodCidr *string `json:"podCIDR,omitempty"`
	// podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node. If this field is specified, the 0th entry must match the podCIDR field. It may contain at most 1 value for each of IPv4 and IPv6.
	PodCidRs []string `json:"podCIDRs,omitempty"`
	// ID of the node assigned by the cloud provider in the format: <ProviderName>://<ProviderSpecificNodeID>
	ProviderID *string `json:"providerID,omitempty"`
	// If specified, the node's taints.
	Taints []TaintV1 `json:"taints,omitempty"`
	// Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
	Unschedulable *bool `json:"unschedulable,omitempty"`
}

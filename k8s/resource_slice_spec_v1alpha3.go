// Code generated; DO NOT EDIT.

package k8s

// ResourceSliceSpecV1alpha3 ResourceSliceSpec contains the information published by the driver in one ResourceSlice.
type ResourceSliceSpecV1alpha3 struct {
	// AllNodes indicates that all nodes have access to the resources in the pool.  Exactly one of NodeName, NodeSelector and AllNodes must be set.
	AllNodes *bool `json:"allNodes,omitempty"`
	// Devices lists some or all of the devices in this pool.  Must not have more than 128 entries.
	Devices []DeviceV1alpha3 `json:"devices,omitempty"`
	// Driver identifies the DRA driver providing the capacity information. A field selector can be used to list only ResourceSlice objects with a certain driver name.  Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver. This field is immutable.
	Driver string `json:"driver"`
	// NodeName identifies the node which provides the resources in this pool. A field selector can be used to list only ResourceSlice objects belonging to a certain node.  This field can be used to limit access from nodes to ResourceSlices with the same node name. It also indicates to autoscalers that adding new nodes of the same type as some old node might also make new resources available.  Exactly one of NodeName, NodeSelector and AllNodes must be set. This field is immutable.
	NodeName *string `json:"nodeName,omitempty"`
	// A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
	NodeSelector *NodeSelectorV1 `json:"nodeSelector,omitempty"`
	// ResourcePool describes the pool that ResourceSlices belong to.
	Pool ResourcePoolV1alpha3 `json:"pool"`
}

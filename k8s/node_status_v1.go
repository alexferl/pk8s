// Code generated; DO NOT EDIT.

package k8s

// NodeStatusV1 NodeStatus is information about the current status of a node.
type NodeStatusV1 struct {
	// List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses Note: This field is declared as mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it is merged. Callers should instead use a full-replacement patch. See https://pr.k8s.io/79391 for an example. Consumers should assume that addresses can change during the lifetime of a Node. However, there are some exceptions where this may not be possible, such as Pods that inherit a Node's address in its own status or consumers of the downward API (status.hostIP).
	Addresses []NodeAddressV1 `json:"addresses,omitempty"`
	// Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
	Allocatable *map[string]string `json:"allocatable,omitempty"`
	// Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/reference/node/node-status/#capacity
	Capacity *map[string]string `json:"capacity,omitempty"`
	// Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition
	Conditions []NodeConditionV1 `json:"conditions,omitempty"`
	// NodeConfigStatus describes the status of the config assigned by Node.Spec.ConfigSource.
	Config *NodeConfigStatusV1 `json:"config,omitempty"`
	// NodeDaemonEndpoints lists ports opened by daemons running on the Node.
	DaemonEndpoints *NodeDaemonEndpointsV1 `json:"daemonEndpoints,omitempty"`
	// NodeFeatures describes the set of features implemented by the CRI implementation. The features contained in the NodeFeatures should depend only on the cri implementation independent of runtime handlers.
	Features *NodeFeaturesV1 `json:"features,omitempty"`
	// List of container images on this node
	Images []ContainerImageV1 `json:"images,omitempty"`
	// NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
	NodeInfo *NodeSystemInfoV1 `json:"nodeInfo,omitempty"`
	// NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
	Phase *string `json:"phase,omitempty"`
	// The available runtime handlers.
	RuntimeHandlers []NodeRuntimeHandlerV1 `json:"runtimeHandlers,omitempty"`
	// List of volumes that are attached to the node.
	VolumesAttached []AttachedVolumeV1 `json:"volumesAttached,omitempty"`
	// List of attachable volumes in use (mounted) by the node.
	VolumesInUse []string `json:"volumesInUse,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// NodeV1 Node is a worker node in Kubernetes. Each node will have a unique identifier in the cache (i.e. in etcd).
type NodeV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// NodeSpec describes the attributes that a node is created with.
	Spec *NodeSpecV1 `json:"spec,omitempty"`
	// NodeStatus is information about the current status of a node.
	Status *NodeStatusV1 `json:"status,omitempty"`
}

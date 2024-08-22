// Code generated; DO NOT EDIT.

package k8s

// DaemonSetV1 DaemonSet represents the configuration of a daemon set.
type DaemonSetV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// DaemonSetSpec is the specification of a daemon set.
	Spec *DaemonSetSpecV1 `json:"spec,omitempty"`
	// DaemonSetStatus represents the current status of a daemon set.
	Status *DaemonSetStatusV1 `json:"status,omitempty"`
}

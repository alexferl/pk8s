// Code generated; DO NOT EDIT.

package k8s

// ServiceV1 Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.
type ServiceV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// ServiceSpec describes the attributes that a user creates on a service.
	Spec *ServiceSpecV1 `json:"spec,omitempty"`
	// ServiceStatus represents the current status of a service.
	Status *ServiceStatusV1 `json:"status,omitempty"`
}

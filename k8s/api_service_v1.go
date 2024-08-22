// Code generated; DO NOT EDIT.

package k8s

// APIServiceV1 APIService represents a server for a particular GroupVersion. Name must be "version.group".
type APIServiceV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.
	Spec *APIServiceSpecV1 `json:"spec,omitempty"`
	// APIServiceStatus contains derived information about an API server
	Status *APIServiceStatusV1 `json:"status,omitempty"`
}

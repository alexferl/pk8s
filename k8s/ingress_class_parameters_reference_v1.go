// Code generated; DO NOT EDIT.

package k8s

// IngressClassParametersReferenceV1 IngressClassParametersReference identifies an API object. This can be used to specify a cluster or namespace-scoped resource.
type IngressClassParametersReferenceV1 struct {
	// apiGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	APIGroup *string `json:"apiGroup,omitempty"`
	// kind is the type of resource being referenced.
	Kind string `json:"kind"`
	// name is the name of resource being referenced.
	Name string `json:"name"`
	// namespace is the namespace of the resource being referenced. This field is required when scope is set to "Namespace" and must be unset when scope is set to "Cluster".
	Namespace *string `json:"namespace,omitempty"`
	// scope represents if this refers to a cluster or namespace scoped resource. This may be set to "Cluster" (default) or "Namespace".
	Scope *string `json:"scope,omitempty"`
}

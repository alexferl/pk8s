// Code generated; DO NOT EDIT.

package k8s

// ParentReferenceV1beta1 ParentReference describes a reference to a parent object.
type ParentReferenceV1beta1 struct {
	// Group is the group of the object being referenced.
	Group *string `json:"group,omitempty"`
	// Name is the name of the object being referenced.
	Name string `json:"name"`
	// Namespace is the namespace of the object being referenced.
	Namespace *string `json:"namespace,omitempty"`
	// Resource is the resource of the object being referenced.
	Resource string `json:"resource"`
}

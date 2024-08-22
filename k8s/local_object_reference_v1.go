// Code generated; DO NOT EDIT.

package k8s

// LocalObjectReferenceV1 LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type LocalObjectReferenceV1 struct {
	// Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`
}

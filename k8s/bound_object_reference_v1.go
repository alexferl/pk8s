// Code generated; DO NOT EDIT.

package k8s

// BoundObjectReferenceV1 BoundObjectReference is a reference to an object that a token is bound to.
type BoundObjectReferenceV1 struct {
	// API version of the referent.
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
	Kind *string `json:"kind,omitempty"`
	// Name of the referent.
	Name *string `json:"name,omitempty"`
	// UID of the referent.
	UID *string `json:"uid,omitempty"`
}

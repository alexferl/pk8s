// Code generated; DO NOT EDIT.

package k8s

// OwnerReferenceV1 OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.
type OwnerReferenceV1 struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion"`
	// If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs "delete" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty"`
	// If true, this reference points to the managing controller.
	Controller *bool `json:"controller,omitempty"`
	// Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	Name string `json:"name"`
	// UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	UID string `json:"uid"`
}

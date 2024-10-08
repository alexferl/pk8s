// Code generated by pk8s; DO NOT EDIT.

package k8s

// ResourceClaimConsumerReferenceV1alpha3 ResourceClaimConsumerReference contains enough information to let you locate the consumer of a ResourceClaim. The user must be a resource in the same namespace as the ResourceClaim.
type ResourceClaimConsumerReferenceV1alpha3 struct {
	// APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
	APIGroup *string `json:"apiGroup,omitempty"`
	// Name is the name of resource being referenced.
	Name string `json:"name"`
	// Resource is the type of resource being referenced, for example "pods".
	Resource string `json:"resource"`
	// UID identifies exactly one incarnation of the resource.
	UID string `json:"uid"`
}

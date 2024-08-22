// Code generated; DO NOT EDIT.

package k8s

// SelfSubjectAccessReviewSpecV1 SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
type SelfSubjectAccessReviewSpecV1 struct {
	// NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	NonResourceAttributes *NonResourceAttributesV1 `json:"nonResourceAttributes,omitempty"`
	// ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
	ResourceAttributes *ResourceAttributesV1 `json:"resourceAttributes,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// TokenReviewV1 TokenReview attempts to authenticate a token to a known user. Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.
type TokenReviewV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// TokenReviewSpec is a description of the token authentication request.
	Spec TokenReviewSpecV1 `json:"spec"`
	// TokenReviewStatus is the result of the token authentication request.
	Status *TokenReviewStatusV1 `json:"status,omitempty"`
}

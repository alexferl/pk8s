// Code generated; DO NOT EDIT.

package k8s

// TokenRequestV1 TokenRequest requests a token for a given service account.
type TokenRequestV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// TokenRequestSpec contains client provided parameters of a token request.
	Spec TokenRequestSpecV1 `json:"spec"`
	// TokenRequestStatus is the result of a token request.
	Status *TokenRequestStatusV1 `json:"status,omitempty"`
}

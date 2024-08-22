// Code generated; DO NOT EDIT.

package k8s

// TokenReviewStatusV1 TokenReviewStatus is the result of the token authentication request.
type TokenReviewStatusV1 struct {
	// Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
	Audiences []string `json:"audiences,omitempty"`
	// Authenticated indicates that the token was associated with a known user.
	Authenticated *bool `json:"authenticated,omitempty"`
	// Error indicates that the token couldn't be checked
	Error *string `json:"error,omitempty"`
	// UserInfo holds the information about the user needed to implement the user.Info interface.
	User *UserInfoV1 `json:"user,omitempty"`
}

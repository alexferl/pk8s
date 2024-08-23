// Code generated by pk8s; DO NOT EDIT.

package k8s

// TokenReviewSpecV1 TokenReviewSpec is a description of the token authentication request.
type TokenReviewSpecV1 struct {
	// Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
	Audiences []string `json:"audiences,omitempty"`
	// Token is the opaque bearer token.
	Token *string `json:"token,omitempty"`
}

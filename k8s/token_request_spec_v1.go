// Code generated; DO NOT EDIT.

package k8s

// TokenRequestSpecV1 TokenRequestSpec contains client provided parameters of a token request.
type TokenRequestSpecV1 struct {
	// Audiences are the intendend audiences of the token. A recipient of a token must identify themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
	Audiences []string `json:"audiences"`
	// BoundObjectReference is a reference to an object that a token is bound to.
	BoundObjectRef *BoundObjectReferenceV1 `json:"boundObjectRef,omitempty"`
	// ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
	ExpirationSeconds *int64 `json:"expirationSeconds,omitempty"`
}

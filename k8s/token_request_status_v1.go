// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// TokenRequestStatusV1 TokenRequestStatus is the result of a token request.
type TokenRequestStatusV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	ExpirationTimestamp time.Time `json:"expirationTimestamp"`
	// Token is the opaque bearer token.
	Token string `json:"token"`
}

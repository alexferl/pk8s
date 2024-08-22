// Code generated; DO NOT EDIT.

package k8s

// SecretKeySelectorV1 SecretKeySelector selects a key of a Secret.
type SecretKeySelectorV1 struct {
	// The key of the secret to select from.  Must be a valid secret key.
	Key string `json:"key"`
	// Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`
	// Specify whether the Secret or its key must be defined
	Optional *bool `json:"optional,omitempty"`
}

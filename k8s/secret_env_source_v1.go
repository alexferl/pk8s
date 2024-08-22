// Code generated; DO NOT EDIT.

package k8s

// SecretEnvSourceV1 SecretEnvSource selects a Secret to populate the environment variables with.  The contents of the target Secret's Data field will represent the key-value pairs as environment variables.
type SecretEnvSourceV1 struct {
	// Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`
	// Specify whether the Secret must be defined
	Optional *bool `json:"optional,omitempty"`
}

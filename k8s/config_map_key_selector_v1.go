// Code generated; DO NOT EDIT.

package k8s

// ConfigMapKeySelectorV1 Selects a key from a ConfigMap.
type ConfigMapKeySelectorV1 struct {
	// The key to select.
	Key string `json:"key"`
	// Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`
	// Specify whether the ConfigMap or its key must be defined
	Optional *bool `json:"optional,omitempty"`
}

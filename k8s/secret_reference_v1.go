// Code generated; DO NOT EDIT.

package k8s

// SecretReferenceV1 SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
type SecretReferenceV1 struct {
	// name is unique within a namespace to reference a secret resource.
	Name *string `json:"name,omitempty"`
	// namespace defines the space within which the secret name must be unique.
	Namespace *string `json:"namespace,omitempty"`
}

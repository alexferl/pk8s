// Code generated; DO NOT EDIT.

package k8s

// EnvFromSourceV1 EnvFromSource represents the source of a set of ConfigMaps
type EnvFromSourceV1 struct {
	// ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.  The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.
	ConfigMapRef *ConfigMapEnvSourceV1 `json:"configMapRef,omitempty"`
	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	Prefix *string `json:"prefix,omitempty"`
	// SecretEnvSource selects a Secret to populate the environment variables with.  The contents of the target Secret's Data field will represent the key-value pairs as environment variables.
	SecretRef *SecretEnvSourceV1 `json:"secretRef,omitempty"`
}

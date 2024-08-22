// Code generated; DO NOT EDIT.

package k8s

// EnvVarSourceV1 EnvVarSource represents a source for the value of an EnvVar.
type EnvVarSourceV1 struct {
	// Selects a key from a ConfigMap.
	ConfigMapKeyRef *ConfigMapKeySelectorV1 `json:"configMapKeyRef,omitempty"`
	// ObjectFieldSelector selects an APIVersioned field of an object.
	FieldRef *ObjectFieldSelectorV1 `json:"fieldRef,omitempty"`
	// ResourceFieldSelector represents container resources (cpu, memory) and their output format
	ResourceFieldRef *ResourceFieldSelectorV1 `json:"resourceFieldRef,omitempty"`
	// SecretKeySelector selects a key of a Secret.
	SecretKeyRef *SecretKeySelectorV1 `json:"secretKeyRef,omitempty"`
}

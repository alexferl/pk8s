// Code generated; DO NOT EDIT.

package k8s

// RollingUpdateDeploymentV1 Spec to control the desired behavior of rolling update.
type RollingUpdateDeploymentV1 struct {
	// IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.
	MaxSurge *string `json:"maxSurge,omitempty"`
	// IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.
	MaxUnavailable *string `json:"maxUnavailable,omitempty"`
}

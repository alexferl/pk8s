// Code generated; DO NOT EDIT.

package k8s

// ScopedResourceSelectorRequirementV1 A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator that relates the scope name and values.
type ScopedResourceSelectorRequirementV1 struct {
	// Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.
	Operator string `json:"operator"`
	// The name of the scope that the selector applies to.
	ScopeName string `json:"scopeName"`
	// An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values []string `json:"values,omitempty"`
}

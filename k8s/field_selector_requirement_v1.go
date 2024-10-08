// Code generated by pk8s; DO NOT EDIT.

package k8s

// FieldSelectorRequirementV1 FieldSelectorRequirement is a selector that contains values, a key, and an operator that relates the key and values.
type FieldSelectorRequirementV1 struct {
	// key is the field selector key that the requirement applies to.
	Key string `json:"key"`
	// operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. The list of operators may grow in the future.
	Operator string `json:"operator"`
	// values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.
	Values []string `json:"values,omitempty"`
}

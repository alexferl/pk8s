// Code generated by pk8s; DO NOT EDIT.

package k8s

// TopologySelectorLabelRequirementV1 A topology selector requirement is a selector that matches given label. This is an alpha feature and may change in the future.
type TopologySelectorLabelRequirementV1 struct {
	// The label key that the selector applies to.
	Key string `json:"key"`
	// An array of string values. One value must match the label to be selected. Each entry in Values is ORed.
	Values []string `json:"values"`
}

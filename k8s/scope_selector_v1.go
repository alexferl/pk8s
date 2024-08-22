// Code generated; DO NOT EDIT.

package k8s

// ScopeSelectorV1 A scope selector represents the AND of the selectors represented by the scoped-resource selector requirements.
type ScopeSelectorV1 struct {
	// A list of scope selector requirements by scope of the resources.
	MatchExpressions []ScopedResourceSelectorRequirementV1 `json:"matchExpressions,omitempty"`
}

// Code generated by pk8s; DO NOT EDIT.

package k8s

// ParamKindV1alpha1 ParamKind is a tuple of Group Kind and Version.
type ParamKindV1alpha1 struct {
	// APIVersion is the API group version the resources belong to. In format of "group/version". Required.
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is the API kind the resources belong to. Required.
	Kind *string `json:"kind,omitempty"`
}

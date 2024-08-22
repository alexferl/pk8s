// Code generated; DO NOT EDIT.

package k8s

// CrossVersionObjectReferenceV2 CrossVersionObjectReference contains enough information to let you identify the referred resource.
type CrossVersionObjectReferenceV2 struct {
	// apiVersion is the API version of the referent
	APIVersion *string `json:"apiVersion,omitempty"`
	// kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind"`
	// name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name"`
}

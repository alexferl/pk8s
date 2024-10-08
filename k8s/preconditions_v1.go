// Code generated by pk8s; DO NOT EDIT.

package k8s

// PreconditionsV1 Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.
type PreconditionsV1 struct {
	// Specifies the target ResourceVersion
	ResourceVersion *string `json:"resourceVersion,omitempty"`
	// Specifies the target UID.
	UID *string `json:"uid,omitempty"`
}

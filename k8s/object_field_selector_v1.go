// Code generated; DO NOT EDIT.

package k8s

// ObjectFieldSelectorV1 ObjectFieldSelector selects an APIVersioned field of an object.
type ObjectFieldSelectorV1 struct {
	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	APIVersion *string `json:"apiVersion,omitempty"`
	// Path of the field to select in the specified API version.
	FieldPath string `json:"fieldPath"`
}

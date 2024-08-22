// Code generated; DO NOT EDIT.

package k8s

// ScaleSpecV1 ScaleSpec describes the attributes of a scale subresource.
type ScaleSpecV1 struct {
	// replicas is the desired number of instances for the scaled object.
	Replicas *int32 `json:"replicas,omitempty"`
}

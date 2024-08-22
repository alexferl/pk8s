// Code generated; DO NOT EDIT.

package k8s

// ScaleStatusV1 ScaleStatus represents the current status of a scale subresource.
type ScaleStatusV1 struct {
	// replicas is the actual number of observed instances of the scaled object.
	Replicas int32 `json:"replicas"`
	// selector is the label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	Selector *string `json:"selector,omitempty"`
}

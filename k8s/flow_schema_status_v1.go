// Code generated; DO NOT EDIT.

package k8s

// FlowSchemaStatusV1 FlowSchemaStatus represents the current state of a FlowSchema.
type FlowSchemaStatusV1 struct {
	// `conditions` is a list of the current states of FlowSchema.
	Conditions []FlowSchemaConditionV1 `json:"conditions,omitempty"`
}

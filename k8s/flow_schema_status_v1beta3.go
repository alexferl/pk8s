// Code generated by pk8s; DO NOT EDIT.

package k8s

// FlowSchemaStatusV1beta3 FlowSchemaStatus represents the current state of a FlowSchema.
type FlowSchemaStatusV1beta3 struct {
	// `conditions` is a list of the current states of FlowSchema.
	Conditions []FlowSchemaConditionV1beta3 `json:"conditions,omitempty"`
}

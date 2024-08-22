// Code generated; DO NOT EDIT.

package k8s

// ComponentConditionV1 Information about the condition of a component.
type ComponentConditionV1 struct {
	// Condition error code for a component. For example, a health check error code.
	Error *string `json:"error,omitempty"`
	// Message about the condition for a component. For example, information about a health check.
	Message *string `json:"message,omitempty"`
	// Status of the condition for a component. Valid values for "Healthy": "True", "False", or "Unknown".
	Status string `json:"status"`
	// Type of condition for a component. Valid value: "Healthy"
	Type string `json:"type"`
}

// Code generated by pk8s; DO NOT EDIT.

package k8s

import (
	"time"
)

// FlowSchemaConditionV1beta3 FlowSchemaCondition describes conditions for a FlowSchema.
type FlowSchemaConditionV1beta3 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// `message` is a human-readable message indicating details about last transition.
	Message *string `json:"message,omitempty"`
	// `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
	Reason *string `json:"reason,omitempty"`
	// `status` is the status of the condition. Can be True, False, Unknown. Required.
	Status *string `json:"status,omitempty"`
	// `type` is the type of the condition. Required.
	Type *string `json:"type,omitempty"`
}

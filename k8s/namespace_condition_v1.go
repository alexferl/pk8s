// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// NamespaceConditionV1 NamespaceCondition contains details about state of namespace.
type NamespaceConditionV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	Message            *string    `json:"message,omitempty"`
	Reason             *string    `json:"reason,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status"`
	// Type of namespace controller condition.
	Type string `json:"type"`
}

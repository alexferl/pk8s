// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// APIServiceConditionV1 APIServiceCondition describes the state of an APIService at a particular point
type APIServiceConditionV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// Human-readable message indicating details about last transition.
	Message *string `json:"message,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason *string `json:"reason,omitempty"`
	// Status is the status of the condition. Can be True, False, Unknown.
	Status string `json:"status"`
	// Type is the type of the condition.
	Type string `json:"type"`
}

// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// CustomResourceDefinitionConditionV1 CustomResourceDefinitionCondition contains details for the current condition of this pod.
type CustomResourceDefinitionConditionV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// message is a human-readable message indicating details about last transition.
	Message *string `json:"message,omitempty"`
	// reason is a unique, one-word, CamelCase reason for the condition's last transition.
	Reason *string `json:"reason,omitempty"`
	// status is the status of the condition. Can be True, False, Unknown.
	Status string `json:"status"`
	// type is the type of the condition. Types include Established, NamesAccepted and Terminating.
	Type string `json:"type"`
}

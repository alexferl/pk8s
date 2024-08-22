// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// NodeConditionV1 NodeCondition contains condition information for a node.
type NodeConditionV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastHeartbeatTime *time.Time `json:"lastHeartbeatTime,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// Human readable message indicating details about last transition.
	Message *string `json:"message,omitempty"`
	// (brief) reason for the condition's last transition.
	Reason *string `json:"reason,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status"`
	// Type of node condition.
	Type string `json:"type"`
}

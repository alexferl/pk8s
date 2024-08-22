// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// ReplicaSetConditionV1 ReplicaSetCondition describes the state of a replica set at a certain point.
type ReplicaSetConditionV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// A human readable message indicating details about the transition.
	Message *string `json:"message,omitempty"`
	// The reason for the condition's last transition.
	Reason *string `json:"reason,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status"`
	// Type of replica set condition.
	Type string `json:"type"`
}

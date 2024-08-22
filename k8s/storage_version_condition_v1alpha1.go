// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// StorageVersionConditionV1alpha1 Describes the state of the storageVersion at a certain point.
type StorageVersionConditionV1alpha1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// A human readable message indicating details about the transition.
	Message string `json:"message"`
	// If set, this represents the .metadata.generation that the condition was set based upon.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason"`
	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status"`
	// Type of the condition.
	Type string `json:"type"`
}

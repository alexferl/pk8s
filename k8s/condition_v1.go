// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// ConditionV1 Condition contains details for one aspect of the current state of this API Resource.
type ConditionV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime time.Time `json:"lastTransitionTime"`
	// message is a human readable message indicating details about the transition. This may be an empty string.
	Message string `json:"message"`
	// observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty.
	Reason string `json:"reason"`
	// status of the condition, one of True, False, Unknown.
	Status string `json:"status"`
	// type of condition in CamelCase or in foo.example.com/CamelCase.
	Type string `json:"type"`
}

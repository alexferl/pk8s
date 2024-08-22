// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// HorizontalPodAutoscalerConditionV2 HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.
type HorizontalPodAutoscalerConditionV2 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// message is a human-readable explanation containing details about the transition
	Message *string `json:"message,omitempty"`
	// reason is the reason for the condition's last transition.
	Reason *string `json:"reason,omitempty"`
	// status is the status of the condition (True, False, Unknown)
	Status string `json:"status"`
	// type describes the current condition
	Type string `json:"type"`
}

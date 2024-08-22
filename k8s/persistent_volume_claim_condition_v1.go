// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// PersistentVolumeClaimConditionV1 PersistentVolumeClaimCondition contains details about state of pvc
type PersistentVolumeClaimConditionV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastProbeTime *time.Time `json:"lastProbeTime,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// message is the human-readable message indicating details about last transition.
	Message *string `json:"message,omitempty"`
	// reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports "Resizing" that means the underlying persistent volume is being resized.
	Reason *string `json:"reason,omitempty"`
	Status string  `json:"status"`
	Type   string  `json:"type"`
}

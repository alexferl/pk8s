// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// PersistentVolumeStatusV1 PersistentVolumeStatus is the current status of a persistent volume.
type PersistentVolumeStatusV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastPhaseTransitionTime *time.Time `json:"lastPhaseTransitionTime,omitempty"`
	// message is a human-readable message indicating details about why the volume is in this state.
	Message *string `json:"message,omitempty"`
	// phase indicates if a volume is available, bound to a claim, or released by a claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase
	Phase *string `json:"phase,omitempty"`
	// reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.
	Reason *string `json:"reason,omitempty"`
}

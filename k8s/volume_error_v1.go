// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// VolumeErrorV1 VolumeError captures an error encountered during a volume operation.
type VolumeErrorV1 struct {
	// message represents the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
	Message *string `json:"message,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	Time *time.Time `json:"time,omitempty"`
}

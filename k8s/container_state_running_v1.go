// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// ContainerStateRunningV1 ContainerStateRunning is a running state of a container.
type ContainerStateRunningV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	StartedAt *time.Time `json:"startedAt,omitempty"`
}

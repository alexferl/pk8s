// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// ContainerStateTerminatedV1 ContainerStateTerminated is a terminated state of a container.
type ContainerStateTerminatedV1 struct {
	// Container's ID in the format '<type>://<container_id>'
	ContainerID *string `json:"containerID,omitempty"`
	// Exit status from the last termination of the container
	ExitCode int32 `json:"exitCode"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// Message regarding the last termination of the container
	Message *string `json:"message,omitempty"`
	// (brief) reason from the last termination of the container
	Reason *string `json:"reason,omitempty"`
	// Signal from the last termination of the container
	Signal *int32 `json:"signal,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	StartedAt *time.Time `json:"startedAt,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// EventSourceV1 EventSource contains information for an event.
type EventSourceV1 struct {
	// Component from which the event is generated.
	Component *string `json:"component,omitempty"`
	// Node name on which the event is generated.
	Host *string `json:"host,omitempty"`
}

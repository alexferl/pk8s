// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// EventV1 Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type EventV1 struct {
	// What action was taken/failed regarding to the Regarding object.
	Action *string `json:"action,omitempty"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// The number of times this event has occurred.
	Count *int32 `json:"count,omitempty"`
	// MicroTime is version of Time with microsecond level precision.
	EventTime *time.Time `json:"eventTime,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	FirstTimestamp *time.Time `json:"firstTimestamp,omitempty"`
	// ObjectReference contains enough information to let you inspect or modify the referred object.
	InvolvedObject ObjectReferenceV1 `json:"involvedObject"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTimestamp *time.Time `json:"lastTimestamp,omitempty"`
	// A human-readable description of the status of this operation.
	Message *string `json:"message,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata ObjectMetaV1 `json:"metadata"`
	// This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
	Reason *string `json:"reason,omitempty"`
	// ObjectReference contains enough information to let you inspect or modify the referred object.
	Related *ObjectReferenceV1 `json:"related,omitempty"`
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingComponent *string `json:"reportingComponent,omitempty"`
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance *string `json:"reportingInstance,omitempty"`
	// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
	Series *EventSeriesV1 `json:"series,omitempty"`
	// EventSource contains information for an event.
	Source *EventSourceV1 `json:"source,omitempty"`
	// Type of this event (Normal, Warning), new types could be added in the future
	Type *string `json:"type,omitempty"`
}

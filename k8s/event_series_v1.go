// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// EventSeriesV1 EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
type EventSeriesV1 struct {
	// Number of occurrences in this series up to the last heartbeat time
	Count *int32 `json:"count,omitempty"`
	// MicroTime is version of Time with microsecond level precision.
	LastObservedTime *time.Time `json:"lastObservedTime,omitempty"`
}

// Code generated by pk8s; DO NOT EDIT.

package k8s

import (
	"time"
)

// TaintV1 The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
type TaintV1 struct {
	// Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	Effect string `json:"effect"`
	// Required. The taint key to be applied to a node.
	Key string `json:"key"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	TimeAdded *time.Time `json:"timeAdded,omitempty"`
	// The taint value corresponding to the taint key.
	Value *string `json:"value,omitempty"`
}

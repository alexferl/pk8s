// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// HorizontalPodAutoscalerStatusV2 HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
type HorizontalPodAutoscalerStatusV2 struct {
	// conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
	Conditions []HorizontalPodAutoscalerConditionV2 `json:"conditions,omitempty"`
	// currentMetrics is the last read state of the metrics used by this autoscaler.
	CurrentMetrics []MetricStatusV2 `json:"currentMetrics,omitempty"`
	// currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
	CurrentReplicas *int32 `json:"currentReplicas,omitempty"`
	// desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
	DesiredReplicas int32 `json:"desiredReplicas"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastScaleTime *time.Time `json:"lastScaleTime,omitempty"`
	// observedGeneration is the most recent generation observed by this autoscaler.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

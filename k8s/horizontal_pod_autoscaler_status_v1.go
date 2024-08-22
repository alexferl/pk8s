// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// HorizontalPodAutoscalerStatusV1 current status of a horizontal pod autoscaler
type HorizontalPodAutoscalerStatusV1 struct {
	// currentCPUUtilizationPercentage is the current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
	CurrentCPUUtilizationPercentage *int32 `json:"currentCPUUtilizationPercentage,omitempty"`
	// currentReplicas is the current number of replicas of pods managed by this autoscaler.
	CurrentReplicas int32 `json:"currentReplicas"`
	// desiredReplicas is the  desired number of replicas of pods managed by this autoscaler.
	DesiredReplicas int32 `json:"desiredReplicas"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastScaleTime *time.Time `json:"lastScaleTime,omitempty"`
	// observedGeneration is the most recent generation observed by this autoscaler.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

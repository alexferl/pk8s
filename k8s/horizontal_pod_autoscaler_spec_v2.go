// Code generated; DO NOT EDIT.

package k8s

// HorizontalPodAutoscalerSpecV2 HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.
type HorizontalPodAutoscalerSpecV2 struct {
	// HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).
	Behavior *HorizontalPodAutoscalerBehaviorV2 `json:"behavior,omitempty"`
	// maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.
	MaxReplicas int32 `json:"maxReplicas"`
	// metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).  The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods.  Ergo, metrics used must decrease as the pod count is increased, and vice-versa.  See the individual metric source types for more information about how each type of metric must respond. If not set, the default metric will be set to 80% average CPU utilization.
	Metrics []MetricSpecV2 `json:"metrics,omitempty"`
	// minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// CrossVersionObjectReference contains enough information to let you identify the referred resource.
	ScaleTargetRef CrossVersionObjectReferenceV2 `json:"scaleTargetRef"`
}

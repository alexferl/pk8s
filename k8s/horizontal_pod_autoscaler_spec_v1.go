// Code generated; DO NOT EDIT.

package k8s

// HorizontalPodAutoscalerSpecV1 specification of a horizontal pod autoscaler.
type HorizontalPodAutoscalerSpecV1 struct {
	// maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
	MaxReplicas int32 `json:"maxReplicas"`
	// minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// CrossVersionObjectReference contains enough information to let you identify the referred resource.
	ScaleTargetRef CrossVersionObjectReferenceV1 `json:"scaleTargetRef"`
	// targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
	TargetCPUUtilizationPercentage *int32 `json:"targetCPUUtilizationPercentage,omitempty"`
}

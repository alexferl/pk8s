// Code generated; DO NOT EDIT.

package k8s

// HorizontalPodAutoscalerBehaviorV2 HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).
type HorizontalPodAutoscalerBehaviorV2 struct {
	// HPAScalingRules configures the scaling behavior for one direction. These Rules are applied after calculating DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead, the safest value from the stabilization window is chosen.
	ScaleDown *HPAScalingRulesV2 `json:"scaleDown,omitempty"`
	// HPAScalingRules configures the scaling behavior for one direction. These Rules are applied after calculating DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead, the safest value from the stabilization window is chosen.
	ScaleUp *HPAScalingRulesV2 `json:"scaleUp,omitempty"`
}

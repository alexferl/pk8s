// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// HorizontalPodAutoscalerBehaviorV2 HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).
type HorizontalPodAutoscalerBehaviorV2 struct {
	ScaleDown *HPAScalingRulesV2 `json:"scaleDown,omitempty"`
	ScaleUp   *HPAScalingRulesV2 `json:"scaleUp,omitempty"`
}

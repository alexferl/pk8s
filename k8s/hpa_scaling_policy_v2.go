// Code generated; DO NOT EDIT.

package k8s

// HPAScalingPolicyV2 HPAScalingPolicy is a single policy which must hold true for a specified past interval.
type HPAScalingPolicyV2 struct {
	// periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).
	PeriodSeconds int32 `json:"periodSeconds"`
	// type is used to specify the scaling policy.
	Type string `json:"type"`
	// value contains the amount of change which is permitted by the policy. It must be greater than zero
	Value int32 `json:"value"`
}

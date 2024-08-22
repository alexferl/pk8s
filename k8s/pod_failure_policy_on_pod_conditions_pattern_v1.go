// Code generated; DO NOT EDIT.

package k8s

// PodFailurePolicyOnPodConditionsPatternV1 PodFailurePolicyOnPodConditionsPattern describes a pattern for matching an actual pod condition type.
type PodFailurePolicyOnPodConditionsPatternV1 struct {
	// Specifies the required Pod condition status. To match a pod condition it is required that the specified status equals the pod condition status. Defaults to True.
	Status string `json:"status"`
	// Specifies the required Pod condition type. To match a pod condition it is required that specified type equals the pod condition type.
	Type string `json:"type"`
}

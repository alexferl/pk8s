// Code generated; DO NOT EDIT.

package k8s

// StatefulSetUpdateStrategyV1 StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
type StatefulSetUpdateStrategyV1 struct {
	// RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.
	RollingUpdate *RollingUpdateStatefulSetStrategyV1 `json:"rollingUpdate,omitempty"`
	// Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
	Type *string `json:"type,omitempty"`
}

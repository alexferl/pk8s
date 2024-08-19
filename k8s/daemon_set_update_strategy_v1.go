// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// DaemonSetUpdateStrategyV1 DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
type DaemonSetUpdateStrategyV1 struct {
	RollingUpdate *RollingUpdateDaemonSetV1 `json:"rollingUpdate,omitempty"`
	// Type of daemon set update. Can be \"RollingUpdate\" or \"OnDelete\". Default is RollingUpdate.
	Type *string `json:"type,omitempty"`
}

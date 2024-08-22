// Code generated; DO NOT EDIT.

package k8s

// DaemonSetUpdateStrategyV1 DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
type DaemonSetUpdateStrategyV1 struct {
	// Spec to control the desired behavior of daemon set rolling update.
	RollingUpdate *RollingUpdateDaemonSetV1 `json:"rollingUpdate,omitempty"`
	// Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
	Type *string `json:"type,omitempty"`
}

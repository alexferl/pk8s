// Code generated; DO NOT EDIT.

package k8s

// DeploymentStrategyV1 DeploymentStrategy describes how to replace existing pods with new ones.
type DeploymentStrategyV1 struct {
	// Spec to control the desired behavior of rolling update.
	RollingUpdate *RollingUpdateDeploymentV1 `json:"rollingUpdate,omitempty"`
	// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	Type *string `json:"type,omitempty"`
}

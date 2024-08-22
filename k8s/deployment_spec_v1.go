// Code generated; DO NOT EDIT.

package k8s

// DeploymentSpecV1 DeploymentSpec is the specification of the desired behavior of the Deployment.
type DeploymentSpecV1 struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	MinReadySeconds *int32 `json:"minReadySeconds,omitempty"`
	// Indicates that the deployment is paused.
	Paused *bool `json:"paused,omitempty"`
	// The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
	ProgressDeadlineSeconds *int32 `json:"progressDeadlineSeconds,omitempty"`
	// Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	Replicas *int32 `json:"replicas,omitempty"`
	// The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Selector LabelSelectorV1 `json:"selector"`
	// DeploymentStrategy describes how to replace existing pods with new ones.
	Strategy *DeploymentStrategyV1 `json:"strategy,omitempty"`
	// PodTemplateSpec describes the data a pod should have when created from a template
	Template PodTemplateSpecV1 `json:"template"`
}

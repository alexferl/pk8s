// Code generated; DO NOT EDIT.

package k8s

// ReplicaSetSpecV1 ReplicaSetSpec is the specification of a ReplicaSet.
type ReplicaSetSpecV1 struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	MinReadySeconds *int32 `json:"minReadySeconds,omitempty"`
	// Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Replicas *int32 `json:"replicas,omitempty"`
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Selector LabelSelectorV1 `json:"selector"`
	// PodTemplateSpec describes the data a pod should have when created from a template
	Template *PodTemplateSpecV1 `json:"template,omitempty"`
}

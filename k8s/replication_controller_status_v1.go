// Code generated; DO NOT EDIT.

package k8s

// ReplicationControllerStatusV1 ReplicationControllerStatus represents the current status of a replication controller.
type ReplicationControllerStatusV1 struct {
	// The number of available replicas (ready for at least minReadySeconds) for this replication controller.
	AvailableReplicas *int32 `json:"availableReplicas,omitempty"`
	// Represents the latest available observations of a replication controller's current state.
	Conditions []ReplicationControllerConditionV1 `json:"conditions,omitempty"`
	// The number of pods that have labels matching the labels of the pod template of the replication controller.
	FullyLabeledReplicas *int32 `json:"fullyLabeledReplicas,omitempty"`
	// ObservedGeneration reflects the generation of the most recently observed replication controller.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// The number of ready replicas for this replication controller.
	ReadyReplicas *int32 `json:"readyReplicas,omitempty"`
	// Replicas is the most recently observed number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Replicas int32 `json:"replicas"`
}

// Code generated; DO NOT EDIT.

package k8s

// ReplicaSetStatusV1 ReplicaSetStatus represents the current status of a ReplicaSet.
type ReplicaSetStatusV1 struct {
	// The number of available replicas (ready for at least minReadySeconds) for this replica set.
	AvailableReplicas *int32 `json:"availableReplicas,omitempty"`
	// Represents the latest available observations of a replica set's current state.
	Conditions []ReplicaSetConditionV1 `json:"conditions,omitempty"`
	// The number of pods that have labels matching the labels of the pod template of the replicaset.
	FullyLabeledReplicas *int32 `json:"fullyLabeledReplicas,omitempty"`
	// ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// readyReplicas is the number of pods targeted by this ReplicaSet with a Ready Condition.
	ReadyReplicas *int32 `json:"readyReplicas,omitempty"`
	// Replicas is the most recently observed number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Replicas int32 `json:"replicas"`
}

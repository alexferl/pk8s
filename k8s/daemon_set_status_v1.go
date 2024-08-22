// Code generated; DO NOT EDIT.

package k8s

// DaemonSetStatusV1 DaemonSetStatus represents the current status of a daemon set.
type DaemonSetStatusV1 struct {
	// Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	CollisionCount *int32 `json:"collisionCount,omitempty"`
	// Represents the latest available observations of a DaemonSet's current state.
	Conditions []DaemonSetConditionV1 `json:"conditions,omitempty"`
	// The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	CurrentNumberScheduled int32 `json:"currentNumberScheduled"`
	// The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	DesiredNumberScheduled int32 `json:"desiredNumberScheduled"`
	// The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)
	NumberAvailable *int32 `json:"numberAvailable,omitempty"`
	// The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	NumberMisscheduled int32 `json:"numberMisscheduled"`
	// numberReady is the number of nodes that should be running the daemon pod and have one or more of the daemon pod running with a Ready Condition.
	NumberReady int32 `json:"numberReady"`
	// The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)
	NumberUnavailable *int32 `json:"numberUnavailable,omitempty"`
	// The most recent generation observed by the daemon set controller.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// The total number of nodes that are running updated daemon pod
	UpdatedNumberScheduled *int32 `json:"updatedNumberScheduled,omitempty"`
}

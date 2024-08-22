// Code generated; DO NOT EDIT.

package k8s

// ContainerStateV1 ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
type ContainerStateV1 struct {
	// ContainerStateRunning is a running state of a container.
	Running *ContainerStateRunningV1 `json:"running,omitempty"`
	// ContainerStateTerminated is a terminated state of a container.
	Terminated *ContainerStateTerminatedV1 `json:"terminated,omitempty"`
	// ContainerStateWaiting is a waiting state of a container.
	Waiting *ContainerStateWaitingV1 `json:"waiting,omitempty"`
}

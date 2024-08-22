// Code generated; DO NOT EDIT.

package k8s

// ContainerStateWaitingV1 ContainerStateWaiting is a waiting state of a container.
type ContainerStateWaitingV1 struct {
	// Message regarding why the container is not yet running.
	Message *string `json:"message,omitempty"`
	// (brief) reason the container is not yet running.
	Reason *string `json:"reason,omitempty"`
}

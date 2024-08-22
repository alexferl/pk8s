// Code generated; DO NOT EDIT.

package k8s

// LifecycleV1 Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.
type LifecycleV1 struct {
	// LifecycleHandler defines a specific action that should be taken in a lifecycle hook. One and only one of the fields, except TCPSocket must be specified.
	PostStart *LifecycleHandlerV1 `json:"postStart,omitempty"`
	// LifecycleHandler defines a specific action that should be taken in a lifecycle hook. One and only one of the fields, except TCPSocket must be specified.
	PreStop *LifecycleHandlerV1 `json:"preStop,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// LifecycleHandlerV1 LifecycleHandler defines a specific action that should be taken in a lifecycle hook. One and only one of the fields, except TCPSocket must be specified.
type LifecycleHandlerV1 struct {
	// ExecAction describes a "run in container" action.
	Exec *ExecActionV1 `json:"exec,omitempty"`
	// HTTPGetAction describes an action based on HTTP Get requests.
	HTTPGet *HTTPGetActionV1 `json:"httpGet,omitempty"`
	// SleepAction describes a "sleep" action.
	Sleep *SleepActionV1 `json:"sleep,omitempty"`
	// TCPSocketAction describes an action based on opening a socket
	TCPSocket *TCPSocketActionV1 `json:"tcpSocket,omitempty"`
}

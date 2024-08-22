// Code generated; DO NOT EDIT.

package k8s

// ContainerPortV1 ContainerPort represents a network port in a single container.
type ContainerPortV1 struct {
	// Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"containerPort"`
	// What host IP to bind the external port to.
	HostIP *string `json:"hostIP,omitempty"`
	// Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
	HostPort *int32 `json:"hostPort,omitempty"`
	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
	Name *string `json:"name,omitempty"`
	// Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
	Protocol *string `json:"protocol,omitempty"`
}

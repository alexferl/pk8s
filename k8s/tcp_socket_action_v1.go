// Code generated; DO NOT EDIT.

package k8s

// TCPSocketActionV1 TCPSocketAction describes an action based on opening a socket
type TCPSocketActionV1 struct {
	// Optional: Host name to connect to, defaults to the pod IP.
	Host *string `json:"host,omitempty"`
	// IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.
	Port string `json:"port"`
}

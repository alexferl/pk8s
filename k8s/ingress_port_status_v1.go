// Code generated; DO NOT EDIT.

package k8s

// IngressPortStatusV1 IngressPortStatus represents the error condition of a service port
type IngressPortStatusV1 struct {
	// error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use   CamelCase names - cloud provider specific error values must have names that comply with the   format foo.example.com/CamelCase.
	Error *string `json:"error,omitempty"`
	// port is the port number of the ingress port.
	Port int32 `json:"port"`
	// protocol is the protocol of the ingress port. The supported values are: "TCP", "UDP", "SCTP"
	Protocol string `json:"protocol"`
}

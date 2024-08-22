// Code generated; DO NOT EDIT.

package k8s

// PortStatusV1
type PortStatusV1 struct {
	// Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use   CamelCase names - cloud provider specific error values must have names that comply with the   format foo.example.com/CamelCase.
	Error *string `json:"error,omitempty"`
	// Port is the port number of the service port of which status is recorded here
	Port int32 `json:"port"`
	// Protocol is the protocol of the service port of which status is recorded here The supported values are: "TCP", "UDP", "SCTP"
	Protocol string `json:"protocol"`
}

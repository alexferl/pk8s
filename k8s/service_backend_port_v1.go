// Code generated; DO NOT EDIT.

package k8s

// ServiceBackendPortV1 ServiceBackendPort is the service port being referenced.
type ServiceBackendPortV1 struct {
	// name is the name of the port on the Service. This is a mutually exclusive setting with "Number".
	Name *string `json:"name,omitempty"`
	// number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with "Name".
	Number *int32 `json:"number,omitempty"`
}

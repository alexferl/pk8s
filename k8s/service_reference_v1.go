// Code generated by pk8s; DO NOT EDIT.

package k8s

// ServiceReferenceV1 ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReferenceV1 struct {
	// `name` is the name of the service. Required
	Name string `json:"name"`
	// `namespace` is the namespace of the service. Required
	Namespace string `json:"namespace"`
	// `path` is an optional URL path which will be sent in any request to this service.
	Path *string `json:"path,omitempty"`
	// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	Port *int32 `json:"port,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// IngressBackendV1 IngressBackend describes all endpoints for a given service and port.
type IngressBackendV1 struct {
	// TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
	Resource *TypedLocalObjectReferenceV1 `json:"resource,omitempty"`
	// IngressServiceBackend references a Kubernetes Service as a Backend.
	Service *IngressServiceBackendV1 `json:"service,omitempty"`
}

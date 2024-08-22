// Code generated; DO NOT EDIT.

package k8s

// IngressServiceBackendV1 IngressServiceBackend references a Kubernetes Service as a Backend.
type IngressServiceBackendV1 struct {
	// name is the referenced service. The service must exist in the same namespace as the Ingress object.
	Name string `json:"name"`
	// ServiceBackendPort is the service port being referenced.
	Port *ServiceBackendPortV1 `json:"port,omitempty"`
}

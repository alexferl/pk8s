// Code generated; DO NOT EDIT.

package k8s

// ServiceStatusV1 ServiceStatus represents the current status of a service.
type ServiceStatusV1 struct {
	// Current service state
	Conditions []ConditionV1 `json:"conditions,omitempty"`
	// LoadBalancerStatus represents the status of a load-balancer.
	LoadBalancer *LoadBalancerStatusV1 `json:"loadBalancer,omitempty"`
}

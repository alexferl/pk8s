// Code generated; DO NOT EDIT.

package k8s

// LoadBalancerStatusV1 LoadBalancerStatus represents the status of a load-balancer.
type LoadBalancerStatusV1 struct {
	// Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.
	Ingress []LoadBalancerIngressV1 `json:"ingress,omitempty"`
}

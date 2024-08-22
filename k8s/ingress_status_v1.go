// Code generated; DO NOT EDIT.

package k8s

// IngressStatusV1 IngressStatus describe the current state of the Ingress.
type IngressStatusV1 struct {
	// IngressLoadBalancerStatus represents the status of a load-balancer.
	LoadBalancer *IngressLoadBalancerStatusV1 `json:"loadBalancer,omitempty"`
}

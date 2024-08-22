// Code generated; DO NOT EDIT.

package k8s

// IngressLoadBalancerStatusV1 IngressLoadBalancerStatus represents the status of a load-balancer.
type IngressLoadBalancerStatusV1 struct {
	// ingress is a list containing ingress points for the load-balancer.
	Ingress []IngressLoadBalancerIngressV1 `json:"ingress,omitempty"`
}

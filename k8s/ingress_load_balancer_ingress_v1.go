// Code generated; DO NOT EDIT.

package k8s

// IngressLoadBalancerIngressV1 IngressLoadBalancerIngress represents the status of a load-balancer ingress point.
type IngressLoadBalancerIngressV1 struct {
	// hostname is set for load-balancer ingress points that are DNS based.
	Hostname *string `json:"hostname,omitempty"`
	// ip is set for load-balancer ingress points that are IP based.
	IP *string `json:"ip,omitempty"`
	// ports provides information about the ports exposed by this LoadBalancer.
	Ports []IngressPortStatusV1 `json:"ports,omitempty"`
}

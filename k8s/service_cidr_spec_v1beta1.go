// Code generated; DO NOT EDIT.

package k8s

// ServiceCIDRSpecV1beta1 ServiceCIDRSpec define the CIDRs the user wants to use for allocating ClusterIPs for Services.
type ServiceCIDRSpecV1beta1 struct {
	// CIDRs defines the IP blocks in CIDR notation (e.g. "192.168.0.0/24" or "2001:db8::/64") from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family. This field is immutable.
	Cidrs []string `json:"cidrs,omitempty"`
}

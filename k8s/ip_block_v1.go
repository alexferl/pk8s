// Code generated; DO NOT EDIT.

package k8s

// IPBlockV1 IPBlock describes a particular CIDR (Ex. "192.168.1.0/24","2001:db8::/64") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
type IPBlockV1 struct {
	// cidr is a string representing the IPBlock Valid examples are "192.168.1.0/24" or "2001:db8::/64"
	Cidr string `json:"cidr"`
	// except is a slice of CIDRs that should not be included within an IPBlock Valid examples are "192.168.1.0/24" or "2001:db8::/64" Except values will be rejected if they are outside the cidr range
	Except []string `json:"except,omitempty"`
}

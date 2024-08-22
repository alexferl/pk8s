// Code generated; DO NOT EDIT.

package k8s

// NodeAddressV1 NodeAddress contains information for the node's address.
type NodeAddressV1 struct {
	// The node address.
	Address string `json:"address"`
	// Node address type, one of Hostname, ExternalIP or InternalIP.
	Type string `json:"type"`
}

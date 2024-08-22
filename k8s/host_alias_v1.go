// Code generated; DO NOT EDIT.

package k8s

// HostAliasV1 HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
type HostAliasV1 struct {
	// Hostnames for the above IP address.
	Hostnames []string `json:"hostnames,omitempty"`
	// IP address of the host file entry.
	IP string `json:"ip"`
}

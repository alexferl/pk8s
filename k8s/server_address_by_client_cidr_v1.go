// Code generated; DO NOT EDIT.

package k8s

// ServerAddressByClientCIDRV1 ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the clientCIDR that they match.
type ServerAddressByClientCIDRV1 struct {
	// The CIDR with which clients can match their IP to figure out the server address that they should use.
	ClientCidr string `json:"clientCIDR"`
	// Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.
	ServerAddress string `json:"serverAddress"`
}

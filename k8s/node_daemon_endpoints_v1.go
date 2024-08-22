// Code generated; DO NOT EDIT.

package k8s

// NodeDaemonEndpointsV1 NodeDaemonEndpoints lists ports opened by daemons running on the Node.
type NodeDaemonEndpointsV1 struct {
	// DaemonEndpoint contains information about a single Daemon endpoint.
	KubeletEndpoint *DaemonEndpointV1 `json:"kubeletEndpoint,omitempty"`
}

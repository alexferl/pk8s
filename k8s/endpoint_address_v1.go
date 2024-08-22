// Code generated; DO NOT EDIT.

package k8s

// EndpointAddressV1 EndpointAddress is a tuple that describes single IP address.
type EndpointAddressV1 struct {
	// The Hostname of this endpoint
	Hostname *string `json:"hostname,omitempty"`
	// The IP of this endpoint. May not be loopback (127.0.0.0/8 or ::1), link-local (169.254.0.0/16 or fe80::/10), or link-local multicast (224.0.0.0/24 or ff02::/16).
	IP string `json:"ip"`
	// Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
	NodeName *string `json:"nodeName,omitempty"`
	// ObjectReference contains enough information to let you inspect or modify the referred object.
	TargetRef *ObjectReferenceV1 `json:"targetRef,omitempty"`
}

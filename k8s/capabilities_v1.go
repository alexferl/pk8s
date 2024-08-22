// Code generated; DO NOT EDIT.

package k8s

// CapabilitiesV1 Adds and removes POSIX capabilities from running containers.
type CapabilitiesV1 struct {
	// Added capabilities
	Add []string `json:"add,omitempty"`
	// Removed capabilities
	Drop []string `json:"drop,omitempty"`
}

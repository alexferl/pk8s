// Code generated; DO NOT EDIT.

package k8s

// ContainerUserV1 ContainerUser represents user identity information
type ContainerUserV1 struct {
	// LinuxContainerUser represents user identity information in Linux containers
	Linux *LinuxContainerUserV1 `json:"linux,omitempty"`
}

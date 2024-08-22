// Code generated; DO NOT EDIT.

package k8s

// SELinuxOptionsV1 SELinuxOptions are the labels to be applied to the container
type SELinuxOptionsV1 struct {
	// Level is SELinux level label that applies to the container.
	Level *string `json:"level,omitempty"`
	// Role is a SELinux role label that applies to the container.
	Role *string `json:"role,omitempty"`
	// Type is a SELinux type label that applies to the container.
	Type *string `json:"type,omitempty"`
	// User is a SELinux user label that applies to the container.
	User *string `json:"user,omitempty"`
}

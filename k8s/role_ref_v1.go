// Code generated; DO NOT EDIT.

package k8s

// RoleRefV1 RoleRef contains information that points to the role being used
type RoleRefV1 struct {
	// APIGroup is the group for the resource being referenced
	APIGroup string `json:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind"`
	// Name is the name of resource being referenced
	Name string `json:"name"`
}

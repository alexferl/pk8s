// Code generated; DO NOT EDIT.

package k8s

// NodeRuntimeHandlerFeaturesV1 NodeRuntimeHandlerFeatures is a set of features implemented by the runtime handler.
type NodeRuntimeHandlerFeaturesV1 struct {
	// RecursiveReadOnlyMounts is set to true if the runtime handler supports RecursiveReadOnlyMounts.
	RecursiveReadOnlyMounts *bool `json:"recursiveReadOnlyMounts,omitempty"`
	// UserNamespaces is set to true if the runtime handler supports UserNamespaces, including for volumes.
	UserNamespaces *bool `json:"userNamespaces,omitempty"`
}

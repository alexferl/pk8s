// Code generated; DO NOT EDIT.

package k8s

// NodeRuntimeHandlerV1 NodeRuntimeHandler is a set of runtime handler information.
type NodeRuntimeHandlerV1 struct {
	// NodeRuntimeHandlerFeatures is a set of features implemented by the runtime handler.
	Features *NodeRuntimeHandlerFeaturesV1 `json:"features,omitempty"`
	// Runtime handler name. Empty for the default runtime handler.
	Name *string `json:"name,omitempty"`
}

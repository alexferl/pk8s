// Code generated; DO NOT EDIT.

package k8s

// NodeFeaturesV1 NodeFeatures describes the set of features implemented by the CRI implementation. The features contained in the NodeFeatures should depend only on the cri implementation independent of runtime handlers.
type NodeFeaturesV1 struct {
	// SupplementalGroupsPolicy is set to true if the runtime supports SupplementalGroupsPolicy and ContainerUser.
	SupplementalGroupsPolicy *bool `json:"supplementalGroupsPolicy,omitempty"`
}

// Code generated by pk8s; DO NOT EDIT.

package k8s

// PodTemplateSpecV1 PodTemplateSpec describes the data a pod should have when created from a template
type PodTemplateSpecV1 struct {
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// PodSpec is a description of a pod.
	Spec *PodSpecV1 `json:"spec,omitempty"`
}

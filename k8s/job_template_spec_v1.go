// Code generated; DO NOT EDIT.

package k8s

// JobTemplateSpecV1 JobTemplateSpec describes the data a Job should have when created from a template
type JobTemplateSpecV1 struct {
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// JobSpec describes how the job execution will look like.
	Spec *JobSpecV1 `json:"spec,omitempty"`
}

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// JobTemplateSpecV1 JobTemplateSpec describes the data a Job should have when created from a template
type JobTemplateSpecV1 struct {
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	Spec     *JobSpecV1    `json:"spec,omitempty"`
}

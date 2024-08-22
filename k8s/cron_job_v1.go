// Code generated; DO NOT EDIT.

package k8s

// CronJobV1 CronJob represents the configuration of a single cron job.
type CronJobV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// CronJobSpec describes how the job execution will look like and when it will actually run.
	Spec *CronJobSpecV1 `json:"spec,omitempty"`
	// CronJobStatus represents the current state of a cron job.
	Status *CronJobStatusV1 `json:"status,omitempty"`
}

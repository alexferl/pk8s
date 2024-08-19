// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

import (
	"time"
)

// CronJobStatusV1 CronJobStatus represents the current state of a cron job.
type CronJobStatusV1 struct {
	// A list of pointers to currently running jobs.
	Active []ObjectReferenceV1 `json:"active,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastScheduleTime *time.Time `json:"lastScheduleTime,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastSuccessfulTime *time.Time `json:"lastSuccessfulTime,omitempty"`
}

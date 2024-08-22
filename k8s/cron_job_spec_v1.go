// Code generated; DO NOT EDIT.

package k8s

// CronJobSpecV1 CronJobSpec describes how the job execution will look like and when it will actually run.
type CronJobSpecV1 struct {
	// Specifies how to treat concurrent executions of a Job. Valid values are:  - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
	ConcurrencyPolicy *string `json:"concurrencyPolicy,omitempty"`
	// The number of failed finished jobs to retain. Value must be non-negative integer. Defaults to 1.
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
	// JobTemplateSpec describes the data a Job should have when created from a template
	JobTemplate JobTemplateSpecV1 `json:"jobTemplate"`
	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Schedule string `json:"schedule"`
	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`
	// The number of successful finished jobs to retain. Value must be non-negative integer. Defaults to 3.
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`
	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
	Suspend *bool `json:"suspend,omitempty"`
	// The time zone name for the given schedule, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. If not specified, this will default to the time zone of the kube-controller-manager process. The set of valid time zone names and the time zone offset is loaded from the system-wide time zone database by the API server during CronJob validation and the controller manager during execution. If no system-wide time zone database can be found a bundled version of the database is used instead. If the time zone name becomes invalid during the lifetime of a CronJob or due to a change in host configuration, the controller will stop creating new new Jobs and will create a system event with the reason UnknownTimeZone. More information can be found in https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones
	TimeZone *string `json:"timeZone,omitempty"`
}

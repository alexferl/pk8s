// Code generated; DO NOT EDIT.

package k8s

// ExternalMetricStatusV2 ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
type ExternalMetricStatusV2 struct {
	// MetricValueStatus holds the current value for a metric
	Current MetricValueStatusV2 `json:"current"`
	// MetricIdentifier defines the name and optionally selector for a metric
	Metric MetricIdentifierV2 `json:"metric"`
}

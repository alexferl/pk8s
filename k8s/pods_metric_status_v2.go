// Code generated; DO NOT EDIT.

package k8s

// PodsMetricStatusV2 PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
type PodsMetricStatusV2 struct {
	// MetricValueStatus holds the current value for a metric
	Current MetricValueStatusV2 `json:"current"`
	// MetricIdentifier defines the name and optionally selector for a metric
	Metric MetricIdentifierV2 `json:"metric"`
}

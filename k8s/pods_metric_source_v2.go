// Code generated; DO NOT EDIT.

package k8s

// PodsMetricSourceV2 PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.
type PodsMetricSourceV2 struct {
	// MetricIdentifier defines the name and optionally selector for a metric
	Metric MetricIdentifierV2 `json:"metric"`
	// MetricTarget defines the target value, average value, or average utilization of a specific metric
	Target MetricTargetV2 `json:"target"`
}

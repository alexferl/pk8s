// Code generated; DO NOT EDIT.

package k8s

// ObjectMetricSourceV2 ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
type ObjectMetricSourceV2 struct {
	// CrossVersionObjectReference contains enough information to let you identify the referred resource.
	DescribedObject CrossVersionObjectReferenceV2 `json:"describedObject"`
	// MetricIdentifier defines the name and optionally selector for a metric
	Metric MetricIdentifierV2 `json:"metric"`
	// MetricTarget defines the target value, average value, or average utilization of a specific metric
	Target MetricTargetV2 `json:"target"`
}

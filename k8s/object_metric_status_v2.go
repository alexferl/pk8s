// Code generated; DO NOT EDIT.

package k8s

// ObjectMetricStatusV2 ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
type ObjectMetricStatusV2 struct {
	// MetricValueStatus holds the current value for a metric
	Current MetricValueStatusV2 `json:"current"`
	// CrossVersionObjectReference contains enough information to let you identify the referred resource.
	DescribedObject CrossVersionObjectReferenceV2 `json:"describedObject"`
	// MetricIdentifier defines the name and optionally selector for a metric
	Metric MetricIdentifierV2 `json:"metric"`
}

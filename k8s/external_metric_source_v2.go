// Code generated; DO NOT EDIT.

package k8s

// ExternalMetricSourceV2 ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
type ExternalMetricSourceV2 struct {
	// MetricIdentifier defines the name and optionally selector for a metric
	Metric MetricIdentifierV2 `json:"metric"`
	// MetricTarget defines the target value, average value, or average utilization of a specific metric
	Target MetricTargetV2 `json:"target"`
}

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// ExternalMetricSourceV2 ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
type ExternalMetricSourceV2 struct {
	Metric MetricIdentifierV2 `json:"metric"`
	Target MetricTargetV2     `json:"target"`
}

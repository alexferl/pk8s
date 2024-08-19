// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// PodsMetricStatusV2 PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
type PodsMetricStatusV2 struct {
	Current MetricValueStatusV2 `json:"current"`
	Metric  MetricIdentifierV2  `json:"metric"`
}

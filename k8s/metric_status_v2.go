// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// MetricStatusV2 MetricStatus describes the last-read state of a single metric.
type MetricStatusV2 struct {
	ContainerResource *ContainerResourceMetricStatusV2 `json:"containerResource,omitempty"`
	External          *ExternalMetricStatusV2          `json:"external,omitempty"`
	Object            *ObjectMetricStatusV2            `json:"object,omitempty"`
	Pods              *PodsMetricStatusV2              `json:"pods,omitempty"`
	Resource          *ResourceMetricStatusV2          `json:"resource,omitempty"`
	// type is the type of metric source.  It will be one of \"ContainerResource\", \"External\", \"Object\", \"Pods\" or \"Resource\", each corresponds to a matching field in the object. Note: \"ContainerResource\" type is available on when the feature-gate HPAContainerMetrics is enabled
	Type string `json:"type"`
}

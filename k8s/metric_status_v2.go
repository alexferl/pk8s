// Code generated; DO NOT EDIT.

package k8s

// MetricStatusV2 MetricStatus describes the last-read state of a single metric.
type MetricStatusV2 struct {
	// ContainerResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing a single container in each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
	ContainerResource *ContainerResourceMetricStatusV2 `json:"containerResource,omitempty"`
	// ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
	External *ExternalMetricStatusV2 `json:"external,omitempty"`
	// ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
	Object *ObjectMetricStatusV2 `json:"object,omitempty"`
	// PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
	Pods *PodsMetricStatusV2 `json:"pods,omitempty"`
	// ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
	Resource *ResourceMetricStatusV2 `json:"resource,omitempty"`
	// type is the type of metric source.  It will be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each corresponds to a matching field in the object. Note: "ContainerResource" type is available on when the feature-gate HPAContainerMetrics is enabled
	Type string `json:"type"`
}

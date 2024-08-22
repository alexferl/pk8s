// Code generated; DO NOT EDIT.

package k8s

// MetricSpecV2 MetricSpec specifies how to scale based on a single metric (only `type` and one other matching field should be set at once).
type MetricSpecV2 struct {
	// ContainerResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
	ContainerResource *ContainerResourceMetricSourceV2 `json:"containerResource,omitempty"`
	// ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
	External *ExternalMetricSourceV2 `json:"external,omitempty"`
	// ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
	Object *ObjectMetricSourceV2 `json:"object,omitempty"`
	// PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.
	Pods *PodsMetricSourceV2 `json:"pods,omitempty"`
	// ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
	Resource *ResourceMetricSourceV2 `json:"resource,omitempty"`
	// type is the type of metric source.  It should be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each mapping to a matching field in the object. Note: "ContainerResource" type is available on when the feature-gate HPAContainerMetrics is enabled
	Type string `json:"type"`
}

// Code generated; DO NOT EDIT.

package k8s

// ResourceMetricSourceV2 ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
type ResourceMetricSourceV2 struct {
	// name is the name of the resource in question.
	Name string `json:"name"`
	// MetricTarget defines the target value, average value, or average utilization of a specific metric
	Target MetricTargetV2 `json:"target"`
}

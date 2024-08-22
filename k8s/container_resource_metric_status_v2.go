// Code generated; DO NOT EDIT.

package k8s

// ContainerResourceMetricStatusV2 ContainerResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing a single container in each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
type ContainerResourceMetricStatusV2 struct {
	// container is the name of the container in the pods of the scaling target
	Container string `json:"container"`
	// MetricValueStatus holds the current value for a metric
	Current MetricValueStatusV2 `json:"current"`
	// name is the name of the resource in question.
	Name string `json:"name"`
}

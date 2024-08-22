// Code generated; DO NOT EDIT.

package k8s

// MetricIdentifierV2 MetricIdentifier defines the name and optionally selector for a metric
type MetricIdentifierV2 struct {
	// name is the name of the given metric
	Name string `json:"name"`
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Selector *LabelSelectorV1 `json:"selector,omitempty"`
}

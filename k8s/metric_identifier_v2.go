// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// MetricIdentifierV2 MetricIdentifier defines the name and optionally selector for a metric
type MetricIdentifierV2 struct {
	// name is the name of the given metric
	Name     string           `json:"name"`
	Selector *LabelSelectorV1 `json:"selector,omitempty"`
}

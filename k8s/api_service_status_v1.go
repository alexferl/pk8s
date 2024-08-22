// Code generated; DO NOT EDIT.

package k8s

// APIServiceStatusV1 APIServiceStatus contains derived information about an API server
type APIServiceStatusV1 struct {
	// Current service state of apiService.
	Conditions []APIServiceConditionV1 `json:"conditions,omitempty"`
}

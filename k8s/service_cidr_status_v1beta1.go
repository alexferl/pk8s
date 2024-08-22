// Code generated; DO NOT EDIT.

package k8s

// ServiceCIDRStatusV1beta1 ServiceCIDRStatus describes the current state of the ServiceCIDR.
type ServiceCIDRStatusV1beta1 struct {
	// conditions holds an array of metav1.Condition that describe the state of the ServiceCIDR. Current service state
	Conditions []ConditionV1 `json:"conditions,omitempty"`
}

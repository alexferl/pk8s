// Code generated; DO NOT EDIT.

package k8s

// ValidatingAdmissionPolicyStatusV1 ValidatingAdmissionPolicyStatus represents the status of an admission validation policy.
type ValidatingAdmissionPolicyStatusV1 struct {
	// The conditions represent the latest available observations of a policy's current state.
	Conditions []ConditionV1 `json:"conditions,omitempty"`
	// The generation observed by the controller.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// TypeChecking contains results of type checking the expressions in the ValidatingAdmissionPolicy
	TypeChecking *TypeCheckingV1 `json:"typeChecking,omitempty"`
}

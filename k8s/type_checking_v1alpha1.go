// Code generated; DO NOT EDIT.

package k8s

// TypeCheckingV1alpha1 TypeChecking contains results of type checking the expressions in the ValidatingAdmissionPolicy
type TypeCheckingV1alpha1 struct {
	// The type checking warnings for each expression.
	ExpressionWarnings []ExpressionWarningV1alpha1 `json:"expressionWarnings,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// TypeCheckingV1beta1 TypeChecking contains results of type checking the expressions in the ValidatingAdmissionPolicy
type TypeCheckingV1beta1 struct {
	// The type checking warnings for each expression.
	ExpressionWarnings []ExpressionWarningV1beta1 `json:"expressionWarnings,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// VariableV1beta1 Variable is the definition of a variable that is used for composition. A variable is defined as a named expression.
type VariableV1beta1 struct {
	// Expression is the expression that will be evaluated as the value of the variable. The CEL expression has access to the same identifiers as the CEL expressions in Validation.
	Expression string `json:"expression"`
	// Name is the name of the variable. The name must be a valid CEL identifier and unique among all variables. The variable can be accessed in other expressions through `variables` For example, if name is "foo", the variable will be available as `variables.foo`
	Name string `json:"name"`
}

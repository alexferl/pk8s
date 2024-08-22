// Code generated; DO NOT EDIT.

package k8s

// EnvVarV1 EnvVar represents an environment variable present in a Container.
type EnvVarV1 struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name"`
	// Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	Value *string `json:"value,omitempty"`
	// EnvVarSource represents a source for the value of an EnvVar.
	ValueFrom *EnvVarSourceV1 `json:"valueFrom,omitempty"`
}

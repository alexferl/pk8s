// Code generated; DO NOT EDIT.

package k8s

// NamespaceStatusV1 NamespaceStatus is information about the current status of a Namespace.
type NamespaceStatusV1 struct {
	// Represents the latest available observations of a namespace's current state.
	Conditions []NamespaceConditionV1 `json:"conditions,omitempty"`
	// Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
	Phase *string `json:"phase,omitempty"`
}

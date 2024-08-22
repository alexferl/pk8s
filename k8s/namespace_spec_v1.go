// Code generated; DO NOT EDIT.

package k8s

// NamespaceSpecV1 NamespaceSpec describes the attributes on a Namespace.
type NamespaceSpecV1 struct {
	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
	Finalizers []string `json:"finalizers,omitempty"`
}

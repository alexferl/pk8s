// Code generated; DO NOT EDIT.

package k8s

// ServiceAccountSubjectV1 ServiceAccountSubject holds detailed information for service-account-kind subject.
type ServiceAccountSubjectV1 struct {
	// `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
	Name string `json:"name"`
	// `namespace` is the namespace of matching ServiceAccount objects. Required.
	Namespace string `json:"namespace"`
}

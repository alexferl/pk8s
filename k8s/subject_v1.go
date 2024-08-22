// Code generated; DO NOT EDIT.

package k8s

// SubjectV1 Subject matches the originator of a request, as identified by the request authentication system. There are three ways of matching an originator; by user, group, or service account.
type SubjectV1 struct {
	// GroupSubject holds detailed information for group-kind subject.
	Group *GroupSubjectV1 `json:"group,omitempty"`
	// `kind` indicates which one of the other fields is non-empty. Required
	Kind string `json:"kind"`
	// ServiceAccountSubject holds detailed information for service-account-kind subject.
	ServiceAccount *ServiceAccountSubjectV1 `json:"serviceAccount,omitempty"`
	// UserSubject holds detailed information for user-kind subject.
	User *UserSubjectV1 `json:"user,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// GroupSubjectV1beta3 GroupSubject holds detailed information for group-kind subject.
type GroupSubjectV1beta3 struct {
	// name is the user group that matches, or "*" to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.
	Name string `json:"name"`
}

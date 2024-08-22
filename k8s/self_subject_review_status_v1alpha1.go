// Code generated; DO NOT EDIT.

package k8s

// SelfSubjectReviewStatusV1alpha1 SelfSubjectReviewStatus is filled by the kube-apiserver and sent back to a user.
type SelfSubjectReviewStatusV1alpha1 struct {
	// UserInfo holds the information about the user needed to implement the user.Info interface.
	UserInfo *UserInfoV1 `json:"userInfo,omitempty"`
}

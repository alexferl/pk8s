// Code generated; DO NOT EDIT.

package k8s

// LinuxContainerUserV1 LinuxContainerUser represents user identity information in Linux containers
type LinuxContainerUserV1 struct {
	// GID is the primary gid initially attached to the first process in the container
	Gid int64 `json:"gid"`
	// SupplementalGroups are the supplemental groups initially attached to the first process in the container
	SupplementalGroups []int64 `json:"supplementalGroups,omitempty"`
	// UID is the primary uid initially attached to the first process in the container
	UID int64 `json:"uid"`
}

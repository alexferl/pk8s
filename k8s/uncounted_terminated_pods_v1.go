// Code generated; DO NOT EDIT.

package k8s

// UncountedTerminatedPodsV1 UncountedTerminatedPods holds UIDs of Pods that have terminated but haven't been accounted in Job status counters.
type UncountedTerminatedPodsV1 struct {
	// failed holds UIDs of failed Pods.
	Failed []string `json:"failed,omitempty"`
	// succeeded holds UIDs of succeeded Pods.
	Succeeded []string `json:"succeeded,omitempty"`
}

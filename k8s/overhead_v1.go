// Code generated; DO NOT EDIT.

package k8s

// OverheadV1 Overhead structure represents the resource overhead associated with running a pod.
type OverheadV1 struct {
	// podFixed represents the fixed resource overhead associated with running a pod.
	PodFixed *map[string]string `json:"podFixed,omitempty"`
}

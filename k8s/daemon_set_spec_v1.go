// Code generated; DO NOT EDIT.

package k8s

// DaemonSetSpecV1 DaemonSetSpec is the specification of a daemon set.
type DaemonSetSpecV1 struct {
	// The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
	MinReadySeconds *int32 `json:"minReadySeconds,omitempty"`
	// The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Selector LabelSelectorV1 `json:"selector"`
	// PodTemplateSpec describes the data a pod should have when created from a template
	Template PodTemplateSpecV1 `json:"template"`
	// DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
	UpdateStrategy *DaemonSetUpdateStrategyV1 `json:"updateStrategy,omitempty"`
}

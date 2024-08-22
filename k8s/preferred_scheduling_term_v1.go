// Code generated; DO NOT EDIT.

package k8s

// PreferredSchedulingTermV1 An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).
type PreferredSchedulingTermV1 struct {
	// A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
	Preference NodeSelectorTermV1 `json:"preference"`
	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
	Weight int32 `json:"weight"`
}

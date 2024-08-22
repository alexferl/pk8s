// Code generated; DO NOT EDIT.

package k8s

// AffinityV1 Affinity is a group of affinity scheduling rules.
type AffinityV1 struct {
	// Node affinity is a group of node affinity scheduling rules.
	NodeAffinity *NodeAffinityV1 `json:"nodeAffinity,omitempty"`
	// Pod affinity is a group of inter pod affinity scheduling rules.
	PodAffinity *PodAffinityV1 `json:"podAffinity,omitempty"`
	// Pod anti affinity is a group of inter pod anti affinity scheduling rules.
	PodAntiAffinity *PodAntiAffinityV1 `json:"podAntiAffinity,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// StatefulSetSpecV1 A StatefulSetSpec is the specification of a StatefulSet.
type StatefulSetSpecV1 struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	MinReadySeconds *int32 `json:"minReadySeconds,omitempty"`
	// StatefulSetOrdinals describes the policy used for replica ordinal assignment in this StatefulSet.
	Ordinals *StatefulSetOrdinalsV1 `json:"ordinals,omitempty"`
	// StatefulSetPersistentVolumeClaimRetentionPolicy describes the policy used for PVCs created from the StatefulSet VolumeClaimTemplates.
	PersistentVolumeClaimRetentionPolicy *StatefulSetPersistentVolumeClaimRetentionPolicyV1 `json:"persistentVolumeClaimRetentionPolicy,omitempty"`
	// podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
	PodManagementPolicy *string `json:"podManagementPolicy,omitempty"`
	// replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
	Replicas *int32 `json:"replicas,omitempty"`
	// revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Selector LabelSelectorV1 `json:"selector"`
	// serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
	ServiceName string `json:"serviceName"`
	// PodTemplateSpec describes the data a pod should have when created from a template
	Template PodTemplateSpecV1 `json:"template"`
	// StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
	UpdateStrategy *StatefulSetUpdateStrategyV1 `json:"updateStrategy,omitempty"`
	// volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
	VolumeClaimTemplates []PersistentVolumeClaimV1 `json:"volumeClaimTemplates,omitempty"`
}

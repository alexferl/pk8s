// Code generated by pk8s; DO NOT EDIT.

package k8s

// PriorityClassV1 PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
type PriorityClassV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description *string `json:"description,omitempty"`
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault *bool `json:"globalDefault,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
	PreemptionPolicy *string `json:"preemptionPolicy,omitempty"`
	// value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value int32 `json:"value"`
}

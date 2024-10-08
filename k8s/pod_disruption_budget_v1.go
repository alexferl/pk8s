// Code generated by pk8s; DO NOT EDIT.

package k8s

// PodDisruptionBudgetV1 PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
type PodDisruptionBudgetV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
	Spec *PodDisruptionBudgetSpecV1 `json:"spec,omitempty"`
	// PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.
	Status *PodDisruptionBudgetStatusV1 `json:"status,omitempty"`
}

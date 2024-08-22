// Code generated; DO NOT EDIT.

package k8s

// PodSchedulingContextV1alpha3 PodSchedulingContext objects hold information that is needed to schedule a Pod with ResourceClaims that use "WaitForFirstConsumer" allocation mode.  This is an alpha type and requires enabling the DRAControlPlaneController feature gate.
type PodSchedulingContextV1alpha3 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// PodSchedulingContextSpec describes where resources for the Pod are needed.
	Spec PodSchedulingContextSpecV1alpha3 `json:"spec"`
	// PodSchedulingContextStatus describes where resources for the Pod can be allocated.
	Status *PodSchedulingContextStatusV1alpha3 `json:"status,omitempty"`
}

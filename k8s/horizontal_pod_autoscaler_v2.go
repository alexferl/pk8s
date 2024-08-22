// Code generated; DO NOT EDIT.

package k8s

// HorizontalPodAutoscalerV2 HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.
type HorizontalPodAutoscalerV2 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.
	Spec *HorizontalPodAutoscalerSpecV2 `json:"spec,omitempty"`
	// HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
	Status *HorizontalPodAutoscalerStatusV2 `json:"status,omitempty"`
}

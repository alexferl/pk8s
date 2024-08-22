// Code generated; DO NOT EDIT.

package k8s

// StorageVersionV1alpha1 Storage version of a specific resource.
type StorageVersionV1alpha1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// StorageVersionSpec is an empty spec.
	Spec map[string]interface{} `json:"spec"`
	// API server instances report the versions they can decode and the version they encode objects to when persisting objects in the backend.
	Status StorageVersionStatusV1alpha1 `json:"status"`
}

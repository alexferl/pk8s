// Code generated; DO NOT EDIT.

package k8s

// PersistentVolumeListV1 PersistentVolumeList is a list of PersistentVolume items.
type PersistentVolumeListV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// items is a list of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
	Items []PersistentVolumeV1 `json:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
	Metadata *ListMetaV1 `json:"metadata,omitempty"`
}

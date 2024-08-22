// Code generated; DO NOT EDIT.

package k8s

// StorageClassV1 StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.  StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
type StorageClassV1 struct {
	// allowVolumeExpansion shows whether the storage class allow volume expand.
	AllowVolumeExpansion *bool `json:"allowVolumeExpansion,omitempty"`
	// allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies []TopologySelectorTermV1 `json:"allowedTopologies,omitempty"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions []string `json:"mountOptions,omitempty"`
	// parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters *map[string]string `json:"parameters,omitempty"`
	// provisioner indicates the type of the provisioner.
	Provisioner string `json:"provisioner"`
	// reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
	ReclaimPolicy *string `json:"reclaimPolicy,omitempty"`
	// volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode *string `json:"volumeBindingMode,omitempty"`
}

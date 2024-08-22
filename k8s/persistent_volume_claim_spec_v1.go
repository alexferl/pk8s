// Code generated; DO NOT EDIT.

package k8s

// PersistentVolumeClaimSpecV1 PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
type PersistentVolumeClaimSpecV1 struct {
	// accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	AccessModes []string `json:"accessModes,omitempty"`
	// TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
	DataSource    *TypedLocalObjectReferenceV1 `json:"dataSource,omitempty"`
	DataSourceRef *TypedObjectReferenceV1      `json:"dataSourceRef,omitempty"`
	// VolumeResourceRequirements describes the storage resource requirements for a volume.
	Resources *VolumeResourceRequirementsV1 `json:"resources,omitempty"`
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Selector *LabelSelectorV1 `json:"selector,omitempty"`
	// storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	StorageClassName *string `json:"storageClassName,omitempty"`
	// volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim. If specified, the CSI driver will create or update the volume with the attributes defined in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName, it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass will be applied to the claim but it's not allowed to reset this field to empty string once it is set. If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass will be set by the persistentvolume controller if it exists. If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource exists. More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/ (Beta) Using this field requires the VolumeAttributesClass feature gate to be enabled (off by default).
	VolumeAttributesClassName *string `json:"volumeAttributesClassName,omitempty"`
	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	VolumeMode *string `json:"volumeMode,omitempty"`
	// volumeName is the binding reference to the PersistentVolume backing this claim.
	VolumeName *string `json:"volumeName,omitempty"`
}

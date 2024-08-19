// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// PersistentVolumeSpecV1 PersistentVolumeSpec is the specification of a persistent volume.
type PersistentVolumeSpecV1 struct {
	// accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	AccessModes          []string                            `json:"accessModes,omitempty"`
	AwsElasticBlockStore *AWSElasticBlockStoreVolumeSourceV1 `json:"awsElasticBlockStore,omitempty"`
	AzureDisk            *AzureDiskVolumeSourceV1            `json:"azureDisk,omitempty"`
	AzureFile            *AzureFilePersistentVolumeSourceV1  `json:"azureFile,omitempty"`
	// capacity is the description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Capacity          *map[string]string                 `json:"capacity,omitempty"`
	Cephfs            *CephFSPersistentVolumeSourceV1    `json:"cephfs,omitempty"`
	Cinder            *CinderPersistentVolumeSourceV1    `json:"cinder,omitempty"`
	ClaimRef          *ObjectReferenceV1                 `json:"claimRef,omitempty"`
	Csi               *CSIPersistentVolumeSourceV1       `json:"csi,omitempty"`
	Fc                *FCVolumeSourceV1                  `json:"fc,omitempty"`
	FlexVolume        *FlexPersistentVolumeSourceV1      `json:"flexVolume,omitempty"`
	Flocker           *FlockerVolumeSourceV1             `json:"flocker,omitempty"`
	GcePersistentDisk *GCEPersistentDiskVolumeSourceV1   `json:"gcePersistentDisk,omitempty"`
	Glusterfs         *GlusterfsPersistentVolumeSourceV1 `json:"glusterfs,omitempty"`
	HostPath          *HostPathVolumeSourceV1            `json:"hostPath,omitempty"`
	Iscsi             *ISCSIPersistentVolumeSourceV1     `json:"iscsi,omitempty"`
	Local             *LocalVolumeSourceV1               `json:"local,omitempty"`
	// mountOptions is the list of mount options, e.g. [\"ro\", \"soft\"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	MountOptions []string              `json:"mountOptions,omitempty"`
	Nfs          *NFSVolumeSourceV1    `json:"nfs,omitempty"`
	NodeAffinity *VolumeNodeAffinityV1 `json:"nodeAffinity,omitempty"`
	// persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	PersistentVolumeReclaimPolicy *string                             `json:"persistentVolumeReclaimPolicy,omitempty"`
	PhotonPersistentDisk          *PhotonPersistentDiskVolumeSourceV1 `json:"photonPersistentDisk,omitempty"`
	PortworxVolume                *PortworxVolumeSourceV1             `json:"portworxVolume,omitempty"`
	Quobyte                       *QuobyteVolumeSourceV1              `json:"quobyte,omitempty"`
	Rbd                           *RBDPersistentVolumeSourceV1        `json:"rbd,omitempty"`
	ScaleIO                       *ScaleIOPersistentVolumeSourceV1    `json:"scaleIO,omitempty"`
	// storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
	StorageClassName *string                            `json:"storageClassName,omitempty"`
	Storageos        *StorageOSPersistentVolumeSourceV1 `json:"storageos,omitempty"`
	// Name of VolumeAttributesClass to which this persistent volume belongs. Empty value is not allowed. When this field is not set, it indicates that this volume does not belong to any VolumeAttributesClass. This field is mutable and can be changed by the CSI driver after a volume has been updated successfully to a new class. For an unbound PersistentVolume, the volumeAttributesClassName will be matched with unbound PersistentVolumeClaims during the binding process. This is an alpha field and requires enabling VolumeAttributesClass feature.
	VolumeAttributesClassName *string `json:"volumeAttributesClassName,omitempty"`
	// volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
	VolumeMode    *string                           `json:"volumeMode,omitempty"`
	VsphereVolume *VsphereVirtualDiskVolumeSourceV1 `json:"vsphereVolume,omitempty"`
}

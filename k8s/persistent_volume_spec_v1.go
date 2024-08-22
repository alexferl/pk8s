// Code generated; DO NOT EDIT.

package k8s

// PersistentVolumeSpecV1 PersistentVolumeSpec is the specification of a persistent volume.
type PersistentVolumeSpecV1 struct {
	// accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	AccessModes []string `json:"accessModes,omitempty"`
	// Represents a Persistent Disk resource in AWS.  An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
	AwsElasticBlockStore *AWSElasticBlockStoreVolumeSourceV1 `json:"awsElasticBlockStore,omitempty"`
	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	AzureDisk *AzureDiskVolumeSourceV1 `json:"azureDisk,omitempty"`
	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	AzureFile *AzureFilePersistentVolumeSourceV1 `json:"azureFile,omitempty"`
	// capacity is the description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Capacity *map[string]string `json:"capacity,omitempty"`
	// Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
	Cephfs *CephFSPersistentVolumeSourceV1 `json:"cephfs,omitempty"`
	// Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
	Cinder *CinderPersistentVolumeSourceV1 `json:"cinder,omitempty"`
	// ObjectReference contains enough information to let you inspect or modify the referred object.
	ClaimRef *ObjectReferenceV1 `json:"claimRef,omitempty"`
	// Represents storage that is managed by an external CSI volume driver (Beta feature)
	Csi *CSIPersistentVolumeSourceV1 `json:"csi,omitempty"`
	// Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
	Fc *FCVolumeSourceV1 `json:"fc,omitempty"`
	// FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.
	FlexVolume *FlexPersistentVolumeSourceV1 `json:"flexVolume,omitempty"`
	// Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
	Flocker *FlockerVolumeSourceV1 `json:"flocker,omitempty"`
	// Represents a Persistent Disk resource in Google Compute Engine.  A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
	GcePersistentDisk *GCEPersistentDiskVolumeSourceV1 `json:"gcePersistentDisk,omitempty"`
	// Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
	Glusterfs *GlusterfsPersistentVolumeSourceV1 `json:"glusterfs,omitempty"`
	// Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
	HostPath *HostPathVolumeSourceV1 `json:"hostPath,omitempty"`
	// ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
	Iscsi *ISCSIPersistentVolumeSourceV1 `json:"iscsi,omitempty"`
	// Local represents directly-attached storage with node affinity (Beta feature)
	Local *LocalVolumeSourceV1 `json:"local,omitempty"`
	// mountOptions is the list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	MountOptions []string `json:"mountOptions,omitempty"`
	// Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
	Nfs *NFSVolumeSourceV1 `json:"nfs,omitempty"`
	// VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
	NodeAffinity *VolumeNodeAffinityV1 `json:"nodeAffinity,omitempty"`
	// persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	PersistentVolumeReclaimPolicy *string `json:"persistentVolumeReclaimPolicy,omitempty"`
	// Represents a Photon Controller persistent disk resource.
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSourceV1 `json:"photonPersistentDisk,omitempty"`
	// PortworxVolumeSource represents a Portworx volume resource.
	PortworxVolume *PortworxVolumeSourceV1 `json:"portworxVolume,omitempty"`
	// Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
	Quobyte *QuobyteVolumeSourceV1 `json:"quobyte,omitempty"`
	// Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
	Rbd *RBDPersistentVolumeSourceV1 `json:"rbd,omitempty"`
	// ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume
	ScaleIo *ScaleIOPersistentVolumeSourceV1 `json:"scaleIO,omitempty"`
	// storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
	StorageClassName *string `json:"storageClassName,omitempty"`
	// Represents a StorageOS persistent volume resource.
	Storageos *StorageOSPersistentVolumeSourceV1 `json:"storageos,omitempty"`
	// Name of VolumeAttributesClass to which this persistent volume belongs. Empty value is not allowed. When this field is not set, it indicates that this volume does not belong to any VolumeAttributesClass. This field is mutable and can be changed by the CSI driver after a volume has been updated successfully to a new class. For an unbound PersistentVolume, the volumeAttributesClassName will be matched with unbound PersistentVolumeClaims during the binding process. This is a beta field and requires enabling VolumeAttributesClass feature (off by default).
	VolumeAttributesClassName *string `json:"volumeAttributesClassName,omitempty"`
	// volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
	VolumeMode *string `json:"volumeMode,omitempty"`
	// Represents a vSphere volume resource.
	VsphereVolume *VsphereVirtualDiskVolumeSourceV1 `json:"vsphereVolume,omitempty"`
}

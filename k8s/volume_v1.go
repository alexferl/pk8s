// Code generated; DO NOT EDIT.

package k8s

// VolumeV1 Volume represents a named volume in a pod that may be accessed by any container in the pod.
type VolumeV1 struct {
	// Represents a Persistent Disk resource in AWS.  An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
	AwsElasticBlockStore *AWSElasticBlockStoreVolumeSourceV1 `json:"awsElasticBlockStore,omitempty"`
	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	AzureDisk *AzureDiskVolumeSourceV1 `json:"azureDisk,omitempty"`
	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	AzureFile *AzureFileVolumeSourceV1 `json:"azureFile,omitempty"`
	// Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
	Cephfs *CephFSVolumeSourceV1 `json:"cephfs,omitempty"`
	// Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
	Cinder *CinderVolumeSourceV1 `json:"cinder,omitempty"`
	// Adapts a ConfigMap into a volume.  The contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.
	ConfigMap *ConfigMapVolumeSourceV1 `json:"configMap,omitempty"`
	// Represents a source location of a volume to mount, managed by an external CSI driver
	Csi *CSIVolumeSourceV1 `json:"csi,omitempty"`
	// DownwardAPIVolumeSource represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.
	DownwardAPI *DownwardAPIVolumeSourceV1 `json:"downwardAPI,omitempty"`
	// Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.
	EmptyDir *EmptyDirVolumeSourceV1 `json:"emptyDir,omitempty"`
	// Represents an ephemeral volume that is handled by a normal storage driver.
	Ephemeral *EphemeralVolumeSourceV1 `json:"ephemeral,omitempty"`
	// Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
	Fc *FCVolumeSourceV1 `json:"fc,omitempty"`
	// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	FlexVolume *FlexVolumeSourceV1 `json:"flexVolume,omitempty"`
	// Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
	Flocker *FlockerVolumeSourceV1 `json:"flocker,omitempty"`
	// Represents a Persistent Disk resource in Google Compute Engine.  A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
	GcePersistentDisk *GCEPersistentDiskVolumeSourceV1 `json:"gcePersistentDisk,omitempty"`
	// Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling.  DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
	GitRepo *GitRepoVolumeSourceV1 `json:"gitRepo,omitempty"`
	// Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
	Glusterfs *GlusterfsVolumeSourceV1 `json:"glusterfs,omitempty"`
	// Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
	HostPath *HostPathVolumeSourceV1 `json:"hostPath,omitempty"`
	// ImageVolumeSource represents a image volume resource.
	Image *ImageVolumeSourceV1 `json:"image,omitempty"`
	// Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
	Iscsi *ISCSIVolumeSourceV1 `json:"iscsi,omitempty"`
	// name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name"`
	// Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
	Nfs *NFSVolumeSourceV1 `json:"nfs,omitempty"`
	// PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSourceV1 `json:"persistentVolumeClaim,omitempty"`
	// Represents a Photon Controller persistent disk resource.
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSourceV1 `json:"photonPersistentDisk,omitempty"`
	// PortworxVolumeSource represents a Portworx volume resource.
	PortworxVolume *PortworxVolumeSourceV1 `json:"portworxVolume,omitempty"`
	// Represents a projected volume source
	Projected *ProjectedVolumeSourceV1 `json:"projected,omitempty"`
	// Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
	Quobyte *QuobyteVolumeSourceV1 `json:"quobyte,omitempty"`
	// Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
	Rbd *RBDVolumeSourceV1 `json:"rbd,omitempty"`
	// ScaleIOVolumeSource represents a persistent ScaleIO volume
	ScaleIo *ScaleIOVolumeSourceV1 `json:"scaleIO,omitempty"`
	// Adapts a Secret into a volume.  The contents of the target Secret's Data field will be presented in a volume as files using the keys in the Data field as the file names. Secret volumes support ownership management and SELinux relabeling.
	Secret *SecretVolumeSourceV1 `json:"secret,omitempty"`
	// Represents a StorageOS persistent volume resource.
	Storageos *StorageOSVolumeSourceV1 `json:"storageos,omitempty"`
	// Represents a vSphere volume resource.
	VsphereVolume *VsphereVirtualDiskVolumeSourceV1 `json:"vsphereVolume,omitempty"`
}

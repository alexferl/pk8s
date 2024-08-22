// Code generated; DO NOT EDIT.

package k8s

// CinderVolumeSourceV1 Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
type CinderVolumeSourceV1 struct {
	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	FsType *string `json:"fsType,omitempty"`
	// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	ReadOnly *bool `json:"readOnly,omitempty"`
	// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	SecretRef *LocalObjectReferenceV1 `json:"secretRef,omitempty"`
	// volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	VolumeID string `json:"volumeID"`
}

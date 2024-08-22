// Code generated; DO NOT EDIT.

package k8s

// RBDVolumeSourceV1 Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
type RBDVolumeSourceV1 struct {
	// fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	FsType *string `json:"fsType,omitempty"`
	// image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Image string `json:"image"`
	// keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Keyring *string `json:"keyring,omitempty"`
	// monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Monitors []string `json:"monitors"`
	// pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Pool *string `json:"pool,omitempty"`
	// readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	ReadOnly *bool `json:"readOnly,omitempty"`
	// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	SecretRef *LocalObjectReferenceV1 `json:"secretRef,omitempty"`
	// user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	User *string `json:"user,omitempty"`
}

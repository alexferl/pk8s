// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// CephFSVolumeSourceV1 Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
type CephFSVolumeSourceV1 struct {
	// monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	Monitors []string `json:"monitors"`
	// path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	Path *string `json:"path,omitempty"`
	// readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	ReadOnly *bool `json:"readOnly,omitempty"`
	// secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	SecretFile *string                 `json:"secretFile,omitempty"`
	SecretRef  *LocalObjectReferenceV1 `json:"secretRef,omitempty"`
	// user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	User *string `json:"user,omitempty"`
}

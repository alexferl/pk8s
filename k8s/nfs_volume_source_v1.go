// Code generated; DO NOT EDIT.

package k8s

// NFSVolumeSourceV1 Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
type NFSVolumeSourceV1 struct {
	// path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Path string `json:"path"`
	// readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	ReadOnly *bool `json:"readOnly,omitempty"`
	// server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Server string `json:"server"`
}

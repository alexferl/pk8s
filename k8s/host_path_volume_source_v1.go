// Code generated; DO NOT EDIT.

package k8s

// HostPathVolumeSourceV1 Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
type HostPathVolumeSourceV1 struct {
	// path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	Path string `json:"path"`
	// type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	Type *string `json:"type,omitempty"`
}

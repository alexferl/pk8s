// Code generated; DO NOT EDIT.

package k8s

// GlusterfsVolumeSourceV1 Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
type GlusterfsVolumeSourceV1 struct {
	// endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Endpoints string `json:"endpoints"`
	// path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Path string `json:"path"`
	// readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	ReadOnly *bool `json:"readOnly,omitempty"`
}

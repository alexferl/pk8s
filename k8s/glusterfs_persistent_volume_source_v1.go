// Code generated; DO NOT EDIT.

package k8s

// GlusterfsPersistentVolumeSourceV1 Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
type GlusterfsPersistentVolumeSourceV1 struct {
	// endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Endpoints string `json:"endpoints"`
	// endpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	EndpointsNamespace *string `json:"endpointsNamespace,omitempty"`
	// path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Path string `json:"path"`
	// readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	ReadOnly *bool `json:"readOnly,omitempty"`
}

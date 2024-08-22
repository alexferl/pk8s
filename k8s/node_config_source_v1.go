// Code generated; DO NOT EDIT.

package k8s

// NodeConfigSourceV1 NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil. This API is deprecated since 1.22
type NodeConfigSourceV1 struct {
	// ConfigMapNodeConfigSource contains the information to reference a ConfigMap as a config source for the Node. This API is deprecated since 1.22: https://git.k8s.io/enhancements/keps/sig-node/281-dynamic-kubelet-configuration
	ConfigMap *ConfigMapNodeConfigSourceV1 `json:"configMap,omitempty"`
}

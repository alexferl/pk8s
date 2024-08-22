// Code generated; DO NOT EDIT.

package k8s

// ConfigMapNodeConfigSourceV1 ConfigMapNodeConfigSource contains the information to reference a ConfigMap as a config source for the Node. This API is deprecated since 1.22: https://git.k8s.io/enhancements/keps/sig-node/281-dynamic-kubelet-configuration
type ConfigMapNodeConfigSourceV1 struct {
	// KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure This field is required in all cases.
	KubeletConfigKey string `json:"kubeletConfigKey"`
	// Name is the metadata.name of the referenced ConfigMap. This field is required in all cases.
	Name string `json:"name"`
	// Namespace is the metadata.namespace of the referenced ConfigMap. This field is required in all cases.
	Namespace string `json:"namespace"`
	// ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
	ResourceVersion *string `json:"resourceVersion,omitempty"`
	// UID is the metadata.UID of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
	UID *string `json:"uid,omitempty"`
}

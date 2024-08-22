// Code generated; DO NOT EDIT.

package k8s

// ContainerImageV1 Describe a container image
type ContainerImageV1 struct {
	// Names by which this image is known. e.g. ["kubernetes.example/hyperkube:v1.0.7", "cloud-vendor.registry.example/cloud-vendor/hyperkube:v1.0.7"]
	Names []string `json:"names,omitempty"`
	// The size of the image in bytes.
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
}

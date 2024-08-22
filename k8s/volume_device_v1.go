// Code generated; DO NOT EDIT.

package k8s

// VolumeDeviceV1 volumeDevice describes a mapping of a raw block device within a container.
type VolumeDeviceV1 struct {
	// devicePath is the path inside of the container that the device will be mapped to.
	DevicePath string `json:"devicePath"`
	// name must match the name of a persistentVolumeClaim in the pod
	Name string `json:"name"`
}

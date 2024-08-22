// Code generated; DO NOT EDIT.

package k8s

// AttachedVolumeV1 AttachedVolume describes a volume attached to a node
type AttachedVolumeV1 struct {
	// DevicePath represents the device path where the volume should be available
	DevicePath string `json:"devicePath"`
	// Name of the attached volume
	Name string `json:"name"`
}

// Code generated; DO NOT EDIT.

package k8s

// DeviceSelectorV1alpha3 DeviceSelector must have exactly one field set.
type DeviceSelectorV1alpha3 struct {
	// CELDeviceSelector contains a CEL expression for selecting a device.
	Cel *CELDeviceSelectorV1alpha3 `json:"cel,omitempty"`
}

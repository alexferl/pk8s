// Code generated by pk8s; DO NOT EDIT.

package k8s

// DeviceClassConfigurationV1alpha3 DeviceClassConfiguration is used in DeviceClass.
type DeviceClassConfigurationV1alpha3 struct {
	// OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
	Opaque *OpaqueDeviceConfigurationV1alpha3 `json:"opaque,omitempty"`
}

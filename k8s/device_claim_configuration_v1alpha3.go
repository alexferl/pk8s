// Code generated; DO NOT EDIT.

package k8s

// DeviceClaimConfigurationV1alpha3 DeviceClaimConfiguration is used for configuration parameters in DeviceClaim.
type DeviceClaimConfigurationV1alpha3 struct {
	// OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
	Opaque *OpaqueDeviceConfigurationV1alpha3 `json:"opaque,omitempty"`
	// Requests lists the names of requests where the configuration applies. If empty, it applies to all requests.
	Requests []string `json:"requests,omitempty"`
}

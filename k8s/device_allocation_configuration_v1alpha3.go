// Code generated by pk8s; DO NOT EDIT.

package k8s

// DeviceAllocationConfigurationV1alpha3 DeviceAllocationConfiguration gets embedded in an AllocationResult.
type DeviceAllocationConfigurationV1alpha3 struct {
	// OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
	Opaque *OpaqueDeviceConfigurationV1alpha3 `json:"opaque,omitempty"`
	// Requests lists the names of requests where the configuration applies. If empty, its applies to all requests.
	Requests []string `json:"requests,omitempty"`
	// Source records whether the configuration comes from a class and thus is not something that a normal user would have been able to set or from a claim.
	Source string `json:"source"`
}

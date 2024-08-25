// Code generated by pk8s; DO NOT EDIT.

package k8s

// DeviceAllocationResultV1alpha3 DeviceAllocationResult is the result of allocating devices.
type DeviceAllocationResultV1alpha3 struct {
	// This field is a combination of all the claim and class configuration parameters. Drivers can distinguish between those based on a flag.  This includes configuration parameters for drivers which have no allocated devices in the result because it is up to the drivers which configuration parameters they support. They can silently ignore unknown configuration parameters.
	Config []DeviceAllocationConfigurationV1alpha3 `json:"config,omitempty"`
	// Results lists all allocated devices.
	Results []DeviceRequestAllocationResultV1alpha3 `json:"results,omitempty"`
}
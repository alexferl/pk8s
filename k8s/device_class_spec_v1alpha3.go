// Code generated; DO NOT EDIT.

package k8s

// DeviceClassSpecV1alpha3 DeviceClassSpec is used in a [DeviceClass] to define what can be allocated and how to configure it.
type DeviceClassSpecV1alpha3 struct {
	// Config defines configuration parameters that apply to each device that is claimed via this class. Some classses may potentially be satisfied by multiple drivers, so each instance of a vendor configuration applies to exactly one driver.  They are passed to the driver, but are not considered while allocating the claim.
	Config []DeviceClassConfigurationV1alpha3 `json:"config,omitempty"`
	// Each selector must be satisfied by a device which is claimed via this class.
	Selectors []DeviceSelectorV1alpha3 `json:"selectors,omitempty"`
	// A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
	SuitableNodes *NodeSelectorV1 `json:"suitableNodes,omitempty"`
}

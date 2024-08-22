// Code generated; DO NOT EDIT.

package k8s

// AllocationResultV1alpha3 AllocationResult contains attributes of an allocated resource.
type AllocationResultV1alpha3 struct {
	// Controller is the name of the DRA driver which handled the allocation. That driver is also responsible for deallocating the claim. It is empty when the claim can be deallocated without involving a driver.  A driver may allocate devices provided by other drivers, so this driver name here can be different from the driver names listed for the results.  This is an alpha field and requires enabling the DRAControlPlaneController feature gate.
	Controller *string `json:"controller,omitempty"`
	// DeviceAllocationResult is the result of allocating devices.
	Devices *DeviceAllocationResultV1alpha3 `json:"devices,omitempty"`
	// A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
	NodeSelector *NodeSelectorV1 `json:"nodeSelector,omitempty"`
}

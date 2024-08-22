// Code generated; DO NOT EDIT.

package k8s

// DeviceClaimV1alpha3 DeviceClaim defines how to request devices with a ResourceClaim.
type DeviceClaimV1alpha3 struct {
	// This field holds configuration for multiple potential drivers which could satisfy requests in this claim. It is ignored while allocating the claim.
	Config []DeviceClaimConfigurationV1alpha3 `json:"config,omitempty"`
	// These constraints must be satisfied by the set of devices that get allocated for the claim.
	Constraints []DeviceConstraintV1alpha3 `json:"constraints,omitempty"`
	// Requests represent individual requests for distinct devices which must all be satisfied. If empty, nothing needs to be allocated.
	Requests []DeviceRequestV1alpha3 `json:"requests,omitempty"`
}

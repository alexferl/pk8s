// Code generated; DO NOT EDIT.

package k8s

// DeviceV1alpha3 Device represents one individual hardware instance that can be selected based on its attributes. Besides the name, exactly one field must be set.
type DeviceV1alpha3 struct {
	// BasicDevice defines one device instance.
	Basic *BasicDeviceV1alpha3 `json:"basic,omitempty"`
	// Name is unique identifier among all devices managed by the driver in the pool. It must be a DNS label.
	Name string `json:"name"`
}

// Code generated; DO NOT EDIT.

package k8s

// BasicDeviceV1alpha3 BasicDevice defines one device instance.
type BasicDeviceV1alpha3 struct {
	// Attributes defines the set of attributes for this device. The name of each attribute must be unique in that set.  The maximum number of attributes and capacities combined is 32.
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	// Capacity defines the set of capacities for this device. The name of each capacity must be unique in that set.  The maximum number of attributes and capacities combined is 32.
	Capacity *map[string]string `json:"capacity,omitempty"`
}

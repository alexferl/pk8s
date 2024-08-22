// Code generated; DO NOT EDIT.

package k8s

// OpaqueDeviceConfigurationV1alpha3 OpaqueDeviceConfiguration contains configuration parameters for a driver in a format defined by the driver vendor.
type OpaqueDeviceConfigurationV1alpha3 struct {
	// Driver is used to determine which kubelet plugin needs to be passed these configuration parameters.  An admission policy provided by the driver developer could use this to decide whether it needs to validate them.  Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
	Driver string `json:"driver"`
	// RawExtension is used to hold extensions in external versions.  To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types.  // Internal package:  	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.Object `json:"myPlugin"` 	}  	type PluginA struct { 		AOption string `json:"aOption"` 	}  // External package:  	type MyAPIObject struct { 		runtime.TypeMeta `json:",inline"` 		MyPlugin runtime.RawExtension `json:"myPlugin"` 	}  	type PluginA struct { 		AOption string `json:"aOption"` 	}  // On the wire, the JSON will look something like this:  	{ 		"kind":"MyAPIObject", 		"apiVersion":"v1", 		"myPlugin": { 			"kind":"PluginA", 			"aOption":"foo", 		}, 	}  So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)
	Parameters map[string]interface{} `json:"parameters"`
}

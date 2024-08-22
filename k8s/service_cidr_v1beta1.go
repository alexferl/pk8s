// Code generated; DO NOT EDIT.

package k8s

// ServiceCIDRV1beta1 ServiceCIDR defines a range of IP addresses using CIDR format (e.g. 192.168.0.0/24 or 2001:db2::/64). This range is used to allocate ClusterIPs to Service objects.
type ServiceCIDRV1beta1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// ServiceCIDRSpec define the CIDRs the user wants to use for allocating ClusterIPs for Services.
	Spec *ServiceCIDRSpecV1beta1 `json:"spec,omitempty"`
	// ServiceCIDRStatus describes the current state of the ServiceCIDR.
	Status *ServiceCIDRStatusV1beta1 `json:"status,omitempty"`
}

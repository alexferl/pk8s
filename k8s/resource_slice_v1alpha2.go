// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// ResourceSliceV1alpha2 ResourceSlice provides information about available resources on individual nodes.
type ResourceSliceV1alpha2 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// DriverName identifies the DRA driver providing the capacity information. A field selector can be used to list only ResourceSlice objects with a certain driver name.
	DriverName string `json:"driverName"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind           *string                          `json:"kind,omitempty"`
	Metadata       *ObjectMetaV1                    `json:"metadata,omitempty"`
	NamedResources *NamedResourcesResourcesV1alpha2 `json:"namedResources,omitempty"`
	// NodeName identifies the node which provides the resources if they are local to a node.  A field selector can be used to list only ResourceSlice objects with a certain node name.
	NodeName *string `json:"nodeName,omitempty"`
}

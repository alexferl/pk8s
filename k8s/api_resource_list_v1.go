// Code generated by pk8s; DO NOT EDIT.

package k8s

// APIResourceListV1 APIResourceList is a list of APIResource, it is used to expose the name of the resources supported in a specific group and version, and if the resource is namespaced.
type APIResourceListV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion *string `json:"apiVersion,omitempty"`
	// groupVersion is the group and version this APIResourceList is for.
	GroupVersion string `json:"groupVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// resources contains the name of the resources and if they are namespaced.
	Resources []APIResourceV1 `json:"resources"`
}

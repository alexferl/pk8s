// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// PodV1 Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.
type PodV1 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string       `json:"kind,omitempty"`
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	Spec     *PodSpecV1    `json:"spec,omitempty"`
	Status   *PodStatusV1  `json:"status,omitempty"`
}

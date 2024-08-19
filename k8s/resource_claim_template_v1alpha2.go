// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// ResourceClaimTemplateV1alpha2 ResourceClaimTemplate is used to produce ResourceClaim objects.
type ResourceClaimTemplateV1alpha2 struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string                           `json:"kind,omitempty"`
	Metadata *ObjectMetaV1                     `json:"metadata,omitempty"`
	Spec     ResourceClaimTemplateSpecV1alpha2 `json:"spec"`
}

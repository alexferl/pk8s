// Code generated; DO NOT EDIT.

package k8s

// ResourceRequirementsV1 ResourceRequirements describes the compute resource requirements.
type ResourceRequirementsV1 struct {
	// Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.
	Claims []ResourceClaimV1 `json:"claims,omitempty"`
	// Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits *map[string]string `json:"limits,omitempty"`
	// Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests *map[string]string `json:"requests,omitempty"`
}

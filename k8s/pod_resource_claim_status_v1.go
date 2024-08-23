// Code generated by pk8s; DO NOT EDIT.

package k8s

// PodResourceClaimStatusV1 PodResourceClaimStatus is stored in the PodStatus for each PodResourceClaim which references a ResourceClaimTemplate. It stores the generated name for the corresponding ResourceClaim.
type PodResourceClaimStatusV1 struct {
	// Name uniquely identifies this resource claim inside the pod. This must match the name of an entry in pod.spec.resourceClaims, which implies that the string must be a DNS_LABEL.
	Name string `json:"name"`
	// ResourceClaimName is the name of the ResourceClaim that was generated for the Pod in the namespace of the Pod. If this is unset, then generating a ResourceClaim was not necessary. The pod.spec.resourceClaims entry can be ignored in this case.
	ResourceClaimName *string `json:"resourceClaimName,omitempty"`
}

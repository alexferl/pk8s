// Code generated; DO NOT EDIT.

package k8s

// PodSchedulingContextStatusV1alpha3 PodSchedulingContextStatus describes where resources for the Pod can be allocated.
type PodSchedulingContextStatusV1alpha3 struct {
	// ResourceClaims describes resource availability for each pod.spec.resourceClaim entry where the corresponding ResourceClaim uses "WaitForFirstConsumer" allocation mode.
	ResourceClaims []ResourceClaimSchedulingStatusV1alpha3 `json:"resourceClaims,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// ResourceClaimV1 ResourceClaim references one entry in PodSpec.ResourceClaims.
type ResourceClaimV1 struct {
	// Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.
	Name string `json:"name"`
	// Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.
	Request *string `json:"request,omitempty"`
}

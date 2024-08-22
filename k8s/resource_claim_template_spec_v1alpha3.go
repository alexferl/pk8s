// Code generated; DO NOT EDIT.

package k8s

// ResourceClaimTemplateSpecV1alpha3 ResourceClaimTemplateSpec contains the metadata and fields for a ResourceClaim.
type ResourceClaimTemplateSpecV1alpha3 struct {
	// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *ObjectMetaV1 `json:"metadata,omitempty"`
	// ResourceClaimSpec defines what is being requested in a ResourceClaim and how to configure it.
	Spec ResourceClaimSpecV1alpha3 `json:"spec"`
}

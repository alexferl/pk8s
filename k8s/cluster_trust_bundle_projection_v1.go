// Code generated by pk8s; DO NOT EDIT.

package k8s

// ClusterTrustBundleProjectionV1 ClusterTrustBundleProjection describes how to select a set of ClusterTrustBundle objects and project their contents into the pod filesystem.
type ClusterTrustBundleProjectionV1 struct {
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	LabelSelector *LabelSelectorV1 `json:"labelSelector,omitempty"`
	// Select a single ClusterTrustBundle by object name.  Mutually-exclusive with signerName and labelSelector.
	Name *string `json:"name,omitempty"`
	// If true, don't block pod startup if the referenced ClusterTrustBundle(s) aren't available.  If using name, then the named ClusterTrustBundle is allowed not to exist.  If using signerName, then the combination of signerName and labelSelector is allowed to match zero ClusterTrustBundles.
	Optional *bool `json:"optional,omitempty"`
	// Relative path from the volume root to write the bundle.
	Path string `json:"path"`
	// Select all ClusterTrustBundles that match this signer name. Mutually-exclusive with name.  The contents of all selected ClusterTrustBundles will be unified and deduplicated.
	SignerName *string `json:"signerName,omitempty"`
}

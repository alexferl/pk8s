// Code generated; DO NOT EDIT.

package k8s

// ResourceQuotaSpecV1 ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
type ResourceQuotaSpecV1 struct {
	// hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Hard *map[string]string `json:"hard,omitempty"`
	// A scope selector represents the AND of the selectors represented by the scoped-resource selector requirements.
	ScopeSelector *ScopeSelectorV1 `json:"scopeSelector,omitempty"`
	// A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
	Scopes []string `json:"scopes,omitempty"`
}

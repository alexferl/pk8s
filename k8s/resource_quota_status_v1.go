// Code generated; DO NOT EDIT.

package k8s

// ResourceQuotaStatusV1 ResourceQuotaStatus defines the enforced hard limits and observed use.
type ResourceQuotaStatusV1 struct {
	// Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Hard *map[string]string `json:"hard,omitempty"`
	// Used is the current observed total usage of the resource in the namespace.
	Used *map[string]string `json:"used,omitempty"`
}

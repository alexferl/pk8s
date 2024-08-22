// Code generated; DO NOT EDIT.

package k8s

// AggregationRuleV1 AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
type AggregationRuleV1 struct {
	// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
	ClusterRoleSelectors []LabelSelectorV1 `json:"clusterRoleSelectors,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// NonResourceRuleV1 NonResourceRule holds information that describes a rule for the non-resource
type NonResourceRuleV1 struct {
	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path.  "*" means all.
	NonResourceUrLs []string `json:"nonResourceURLs,omitempty"`
	// Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options.  "*" means all.
	Verbs []string `json:"verbs"`
}

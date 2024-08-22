// Code generated; DO NOT EDIT.

package k8s

// ResourceAttributesV1 ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
type ResourceAttributesV1 struct {
	// FieldSelectorAttributes indicates a field limited access. Webhook authors are encouraged to * ensure rawSelector and requirements are not both set * consider the requirements field if set * not try to parse or consider the rawSelector field if set. This is to avoid another CVE-2022-2880 (i.e. getting different systems to agree on how exactly to parse a query is not something we want), see https://www.oxeye.io/resources/golang-parameter-smuggling-attack for more details. For the *SubjectAccessReview endpoints of the kube-apiserver: * If rawSelector is empty and requirements are empty, the request is not limited. * If rawSelector is present and requirements are empty, the rawSelector will be parsed and limited if the parsing succeeds. * If rawSelector is empty and requirements are present, the requirements should be honored * If rawSelector is present and requirements are present, the request is invalid.
	FieldSelector *FieldSelectorAttributesV1 `json:"fieldSelector,omitempty"`
	// Group is the API Group of the Resource.  "*" means all.
	Group *string `json:"group,omitempty"`
	// LabelSelectorAttributes indicates a label limited access. Webhook authors are encouraged to * ensure rawSelector and requirements are not both set * consider the requirements field if set * not try to parse or consider the rawSelector field if set. This is to avoid another CVE-2022-2880 (i.e. getting different systems to agree on how exactly to parse a query is not something we want), see https://www.oxeye.io/resources/golang-parameter-smuggling-attack for more details. For the *SubjectAccessReview endpoints of the kube-apiserver: * If rawSelector is empty and requirements are empty, the request is not limited. * If rawSelector is present and requirements are empty, the rawSelector will be parsed and limited if the parsing succeeds. * If rawSelector is empty and requirements are present, the requirements should be honored * If rawSelector is present and requirements are present, the request is invalid.
	LabelSelector *LabelSelectorAttributesV1 `json:"labelSelector,omitempty"`
	// Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
	Name *string `json:"name,omitempty"`
	// Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	Namespace *string `json:"namespace,omitempty"`
	// Resource is one of the existing resource types.  "*" means all.
	Resource *string `json:"resource,omitempty"`
	// Subresource is one of the existing resource types.  "" means none.
	Subresource *string `json:"subresource,omitempty"`
	// Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.
	Verb *string `json:"verb,omitempty"`
	// Version is the API Version of the Resource.  "*" means all.
	Version *string `json:"version,omitempty"`
}

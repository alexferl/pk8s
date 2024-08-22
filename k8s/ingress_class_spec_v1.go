// Code generated; DO NOT EDIT.

package k8s

// IngressClassSpecV1 IngressClassSpec provides information about the class of an Ingress.
type IngressClassSpecV1 struct {
	// controller refers to the name of the controller that should handle this class. This allows for different "flavors" that are controlled by the same controller. For example, you may have different parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. "acme.io/ingress-controller". This field is immutable.
	Controller *string `json:"controller,omitempty"`
	// IngressClassParametersReference identifies an API object. This can be used to specify a cluster or namespace-scoped resource.
	Parameters *IngressClassParametersReferenceV1 `json:"parameters,omitempty"`
}

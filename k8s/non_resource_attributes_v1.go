// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// NonResourceAttributesV1 NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
type NonResourceAttributesV1 struct {
	// Path is the URL path of the request
	Path *string `json:"path,omitempty"`
	// Verb is the standard HTTP verb
	Verb *string `json:"verb,omitempty"`
}

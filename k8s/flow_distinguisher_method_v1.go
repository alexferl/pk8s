// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// FlowDistinguisherMethodV1 FlowDistinguisherMethod specifies the method of a flow distinguisher.
type FlowDistinguisherMethodV1 struct {
	// `type` is the type of flow distinguisher method The supported types are \"ByUser\" and \"ByNamespace\". Required.
	Type string `json:"type"`
}

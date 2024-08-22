// Code generated; DO NOT EDIT.

package k8s

// FlowDistinguisherMethodV1beta3 FlowDistinguisherMethod specifies the method of a flow distinguisher.
type FlowDistinguisherMethodV1beta3 struct {
	// `type` is the type of flow distinguisher method The supported types are "ByUser" and "ByNamespace". Required.
	Type string `json:"type"`
}

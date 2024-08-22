// Code generated; DO NOT EDIT.

package k8s

// ResourceStatusV1
type ResourceStatusV1 struct {
	// Name of the resource. Must be unique within the pod and match one of the resources from the pod spec.
	Name string `json:"name"`
	// List of unique Resources health. Each element in the list contains an unique resource ID and resource health. At a minimum, ResourceID must uniquely identify the Resource allocated to the Pod on the Node for the lifetime of a Pod. See ResourceID type for it's definition.
	Resources []ResourceHealthV1 `json:"resources,omitempty"`
}

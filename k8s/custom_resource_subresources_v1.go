// Code generated; DO NOT EDIT.

package k8s

// CustomResourceSubresourcesV1 CustomResourceSubresources defines the status and scale subresources for CustomResources.
type CustomResourceSubresourcesV1 struct {
	// CustomResourceSubresourceScale defines how to serve the scale subresource for CustomResources.
	Scale *CustomResourceSubresourceScaleV1 `json:"scale,omitempty"`
	// CustomResourceSubresourceStatus defines how to serve the status subresource for CustomResources. Status is represented by the `.status` JSON path inside of a CustomResource. When set, * exposes a /status subresource for the custom resource * PUT requests to the /status subresource take a custom resource object, and ignore changes to anything except the status stanza * PUT/POST/PATCH requests to the custom resource ignore changes to the status stanza
	Status map[string]interface{} `json:"status,omitempty"`
}

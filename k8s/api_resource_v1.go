// Code generated; DO NOT EDIT.

package k8s

// APIResourceV1 APIResource specifies the name of a resource and whether it is namespaced.
type APIResourceV1 struct {
	// categories is a list of the grouped resources this resource belongs to (e.g. 'all')
	Categories []string `json:"categories,omitempty"`
	// group is the preferred group of the resource.  Empty implies the group of the containing resource list. For subresources, this may have a different value, for example: Scale".
	Group *string `json:"group,omitempty"`
	// kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')
	Kind string `json:"kind"`
	// name is the plural name of the resource.
	Name string `json:"name"`
	// namespaced indicates if a resource is namespaced or not.
	Namespaced bool `json:"namespaced"`
	// shortNames is a list of suggested short names of the resource.
	ShortNames []string `json:"shortNames,omitempty"`
	// singularName is the singular name of the resource.  This allows clients to handle plural and singular opaquely. The singularName is more correct for reporting status on a single item and both singular and plural are allowed from the kubectl CLI interface.
	SingularName string `json:"singularName"`
	// The hash value of the storage version, the version this resource is converted to when written to the data store. Value must be treated as opaque by clients. Only equality comparison on the value is valid. This is an alpha feature and may change or be removed in the future. The field is populated by the apiserver only if the StorageVersionHash feature gate is enabled. This field will remain optional even if it graduates.
	StorageVersionHash *string `json:"storageVersionHash,omitempty"`
	// verbs is a list of supported kube verbs (this includes get, list, watch, create, update, patch, delete, deletecollection, and proxy)
	Verbs []string `json:"verbs"`
	// version is the preferred version of the resource.  Empty implies the version of the containing resource list For subresources, this may have a different value, for example: v1 (while inside a v1beta1 version of the core resource's group)".
	Version *string `json:"version,omitempty"`
}

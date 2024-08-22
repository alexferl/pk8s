// Code generated; DO NOT EDIT.

package k8s

// CustomResourceDefinitionStatusV1 CustomResourceDefinitionStatus indicates the state of the CustomResourceDefinition
type CustomResourceDefinitionStatusV1 struct {
	// CustomResourceDefinitionNames indicates the names to serve this CustomResourceDefinition
	AcceptedNames *CustomResourceDefinitionNamesV1 `json:"acceptedNames,omitempty"`
	// conditions indicate state for particular aspects of a CustomResourceDefinition
	Conditions []CustomResourceDefinitionConditionV1 `json:"conditions,omitempty"`
	// storedVersions lists all versions of CustomResources that were ever persisted. Tracking these versions allows a migration path for stored versions in etcd. The field is mutable so a migration controller can finish a migration to another version (ensuring no old objects are left in storage), and then remove the rest of the versions from this list. Versions may not be removed from `spec.versions` while they exist in this list.
	StoredVersions []string `json:"storedVersions,omitempty"`
}

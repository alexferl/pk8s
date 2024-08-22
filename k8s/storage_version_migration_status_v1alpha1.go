// Code generated; DO NOT EDIT.

package k8s

// StorageVersionMigrationStatusV1alpha1 Status of the storage version migration.
type StorageVersionMigrationStatusV1alpha1 struct {
	// The latest available observations of the migration's current state.
	Conditions []MigrationConditionV1alpha1 `json:"conditions,omitempty"`
	// ResourceVersion to compare with the GC cache for performing the migration. This is the current resource version of given group, version and resource when kube-controller-manager first observes this StorageVersionMigration resource.
	ResourceVersion *string `json:"resourceVersion,omitempty"`
}

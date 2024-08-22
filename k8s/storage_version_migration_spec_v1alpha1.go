// Code generated; DO NOT EDIT.

package k8s

// StorageVersionMigrationSpecV1alpha1 Spec of the storage version migration.
type StorageVersionMigrationSpecV1alpha1 struct {
	// The token used in the list options to get the next chunk of objects to migrate. When the .status.conditions indicates the migration is "Running", users can use this token to check the progress of the migration.
	ContinueToken *string `json:"continueToken,omitempty"`
	// The names of the group, the version, and the resource.
	Resource GroupVersionResourceV1alpha1 `json:"resource"`
}

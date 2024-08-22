// Code generated; DO NOT EDIT.

package k8s

// StorageVersionStatusV1alpha1 API server instances report the versions they can decode and the version they encode objects to when persisting objects in the backend.
type StorageVersionStatusV1alpha1 struct {
	// If all API server instances agree on the same encoding storage version, then this field is set to that version. Otherwise this field is left empty. API servers should finish updating its storageVersionStatus entry before serving write operations, so that this field will be in sync with the reality.
	CommonEncodingVersion *string `json:"commonEncodingVersion,omitempty"`
	// The latest available observations of the storageVersion's state.
	Conditions []StorageVersionConditionV1alpha1 `json:"conditions,omitempty"`
	// The reported versions per API server instance.
	StorageVersions []ServerStorageVersionV1alpha1 `json:"storageVersions,omitempty"`
}

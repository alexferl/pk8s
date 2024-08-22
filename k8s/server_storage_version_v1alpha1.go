// Code generated; DO NOT EDIT.

package k8s

// ServerStorageVersionV1alpha1 An API server instance reports the version it can decode and the version it encodes objects to when persisting objects in the backend.
type ServerStorageVersionV1alpha1 struct {
	// The ID of the reporting API server.
	APIServerID *string `json:"apiServerID,omitempty"`
	// The API server can decode objects encoded in these versions. The encodingVersion must be included in the decodableVersions.
	DecodableVersions []string `json:"decodableVersions,omitempty"`
	// The API server encodes the object to this version when persisting it in the backend (e.g., etcd).
	EncodingVersion *string `json:"encodingVersion,omitempty"`
	// The API server can serve these versions. DecodableVersions must include all ServedVersions.
	ServedVersions []string `json:"servedVersions,omitempty"`
}

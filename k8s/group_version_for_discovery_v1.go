// Code generated; DO NOT EDIT.

package k8s

// GroupVersionForDiscoveryV1 GroupVersion contains the "group/version" and "version" string of a version. It is made a struct to keep extensibility.
type GroupVersionForDiscoveryV1 struct {
	// groupVersion specifies the API group and version in the form "group/version"
	GroupVersion string `json:"groupVersion"`
	// version specifies the version in the form of "version". This is to save the clients the trouble of splitting the GroupVersion.
	Version string `json:"version"`
}

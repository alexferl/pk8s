// Code generated; DO NOT EDIT.

package k8s

// DeviceAttributeV1alpha3 DeviceAttribute must have exactly one field set.
type DeviceAttributeV1alpha3 struct {
	// BoolValue is a true/false value.
	Bool *bool `json:"bool,omitempty"`
	// IntValue is a number.
	Int *int64 `json:"int,omitempty"`
	// StringValue is a string. Must not be longer than 64 characters.
	String *string `json:"string,omitempty"`
	// VersionValue is a semantic version according to semver.org spec 2.0.0. Must not be longer than 64 characters.
	Version *string `json:"version,omitempty"`
}

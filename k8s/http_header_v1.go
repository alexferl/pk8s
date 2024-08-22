// Code generated; DO NOT EDIT.

package k8s

// HTTPHeaderV1 HTTPHeader describes a custom header to be used in HTTP probes
type HTTPHeaderV1 struct {
	// The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.
	Name string `json:"name"`
	// The header field value
	Value string `json:"value"`
}

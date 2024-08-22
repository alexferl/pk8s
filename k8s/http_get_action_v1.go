// Code generated; DO NOT EDIT.

package k8s

// HTTPGetActionV1 HTTPGetAction describes an action based on HTTP Get requests.
type HTTPGetActionV1 struct {
	// Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
	Host *string `json:"host,omitempty"`
	// Custom headers to set in the request. HTTP allows repeated headers.
	HTTPHeaders []HTTPHeaderV1 `json:"httpHeaders,omitempty"`
	// Path to access on the HTTP server.
	Path *string `json:"path,omitempty"`
	// IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.
	Port string `json:"port"`
	// Scheme to use for connecting to the host. Defaults to HTTP.
	Scheme *string `json:"scheme,omitempty"`
}

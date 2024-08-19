// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s

// LimitResponseV1beta3 LimitResponse defines how to handle requests that can not be executed right now.
type LimitResponseV1beta3 struct {
	Queuing *QueuingConfigurationV1beta3 `json:"queuing,omitempty"`
	// `type` is \"Queue\" or \"Reject\". \"Queue\" means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. \"Reject\" means that requests that can not be executed upon arrival are rejected. Required.
	Type string `json:"type"`
}

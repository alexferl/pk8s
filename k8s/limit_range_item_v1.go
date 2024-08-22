// Code generated; DO NOT EDIT.

package k8s

// LimitRangeItemV1 LimitRangeItem defines a min/max usage limit for any resource that matches on kind.
type LimitRangeItemV1 struct {
	// Default resource requirement limit value by resource name if resource limit is omitted.
	Default *map[string]string `json:"default,omitempty"`
	// DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
	DefaultRequest *map[string]string `json:"defaultRequest,omitempty"`
	// Max usage constraints on this kind by resource name.
	Max *map[string]string `json:"max,omitempty"`
	// MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
	MaxLimitRequestRatio *map[string]string `json:"maxLimitRequestRatio,omitempty"`
	// Min usage constraints on this kind by resource name.
	Min *map[string]string `json:"min,omitempty"`
	// Type of resource that this limit applies to.
	Type string `json:"type"`
}

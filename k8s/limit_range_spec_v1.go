// Code generated; DO NOT EDIT.

package k8s

// LimitRangeSpecV1 LimitRangeSpec defines a min/max usage limit for resources that match on kind.
type LimitRangeSpecV1 struct {
	// Limits is the list of LimitRangeItem objects that are enforced.
	Limits []LimitRangeItemV1 `json:"limits"`
}

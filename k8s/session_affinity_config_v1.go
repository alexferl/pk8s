// Code generated; DO NOT EDIT.

package k8s

// SessionAffinityConfigV1 SessionAffinityConfig represents the configurations of session affinity.
type SessionAffinityConfigV1 struct {
	// ClientIPConfig represents the configurations of Client IP based session affinity.
	ClientIP *ClientIPConfigV1 `json:"clientIP,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// ClientIPConfigV1 ClientIPConfig represents the configurations of Client IP based session affinity.
type ClientIPConfigV1 struct {
	// timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP". Default value is 10800(for 3 hours).
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

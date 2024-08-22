// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// CertificateSigningRequestConditionV1 CertificateSigningRequestCondition describes a condition of a CertificateSigningRequest object
type CertificateSigningRequestConditionV1 struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	// message contains a human readable message with details about the request state
	Message *string `json:"message,omitempty"`
	// reason indicates a brief reason for the request state
	Reason *string `json:"reason,omitempty"`
	// status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown".
	Status string `json:"status"`
	// type of the condition. Known conditions are "Approved", "Denied", and "Failed".  An "Approved" condition is added via the /approval subresource, indicating the request was approved and should be issued by the signer.  A "Denied" condition is added via the /approval subresource, indicating the request was denied and should not be issued by the signer.  A "Failed" condition is added via the /status subresource, indicating the signer failed to issue the certificate.  Approved and Denied conditions are mutually exclusive. Approved, Denied, and Failed conditions cannot be removed once added.  Only one condition of a given type is allowed.
	Type string `json:"type"`
}

// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// LeaseSpecV1 LeaseSpec is a specification of a Lease.
type LeaseSpecV1 struct {
	// MicroTime is version of Time with microsecond level precision.
	AcquireTime *time.Time `json:"acquireTime,omitempty"`
	// holderIdentity contains the identity of the holder of a current lease. If Coordinated Leader Election is used, the holder identity must be equal to the elected LeaseCandidate.metadata.name field.
	HolderIdentity *string `json:"holderIdentity,omitempty"`
	// leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measured against the time of last observed renewTime.
	LeaseDurationSeconds *int32 `json:"leaseDurationSeconds,omitempty"`
	// leaseTransitions is the number of transitions of a lease between holders.
	LeaseTransitions *int32 `json:"leaseTransitions,omitempty"`
	// PreferredHolder signals to a lease holder that the lease has a more optimal holder and should be given up. This field can only be set if Strategy is also set.
	PreferredHolder *string `json:"preferredHolder,omitempty"`
	// MicroTime is version of Time with microsecond level precision.
	RenewTime *time.Time `json:"renewTime,omitempty"`
	// Strategy indicates the strategy for picking the leader for coordinated leader election. If the field is not specified, there is no active coordination for this lease. (Alpha) Using this field requires the CoordinatedLeaderElection feature gate to be enabled.
	Strategy *string `json:"strategy,omitempty"`
}

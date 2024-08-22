// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// LeaseCandidateSpecV1alpha1 LeaseCandidateSpec is a specification of a Lease.
type LeaseCandidateSpecV1alpha1 struct {
	// BinaryVersion is the binary version. It must be in a semver format without leading `v`. This field is required when strategy is "OldestEmulationVersion"
	BinaryVersion *string `json:"binaryVersion,omitempty"`
	// EmulationVersion is the emulation version. It must be in a semver format without leading `v`. EmulationVersion must be less than or equal to BinaryVersion. This field is required when strategy is "OldestEmulationVersion"
	EmulationVersion *string `json:"emulationVersion,omitempty"`
	// LeaseName is the name of the lease for which this candidate is contending. This field is immutable.
	LeaseName string `json:"leaseName"`
	// MicroTime is version of Time with microsecond level precision.
	PingTime *time.Time `json:"pingTime,omitempty"`
	// PreferredStrategies indicates the list of strategies for picking the leader for coordinated leader election. The list is ordered, and the first strategy supersedes all other strategies. The list is used by coordinated leader election to make a decision about the final election strategy. This follows as - If all clients have strategy X as the first element in this list, strategy X will be used. - If a candidate has strategy [X] and another candidate has strategy [Y, X], Y supersedes X and strategy Y   will be used. - If a candidate has strategy [X, Y] and another candidate has strategy [Y, X], this is a user error and leader   election will not operate the Lease until resolved. (Alpha) Using this field requires the CoordinatedLeaderElection feature gate to be enabled.
	PreferredStrategies []string `json:"preferredStrategies"`
	// MicroTime is version of Time with microsecond level precision.
	RenewTime *time.Time `json:"renewTime,omitempty"`
}

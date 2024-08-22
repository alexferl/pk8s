// Code generated; DO NOT EDIT.

package k8s

// StatusDetailsV1 StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.
type StatusDetailsV1 struct {
	// The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.
	Causes []StatusCauseV1 `json:"causes,omitempty"`
	// The group attribute of the resource associated with the status StatusReason.
	Group *string `json:"group,omitempty"`
	// The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).
	Name *string `json:"name,omitempty"`
	// If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.
	RetryAfterSeconds *int32 `json:"retryAfterSeconds,omitempty"`
	// UID of the resource. (when there is a single resource which can be described). More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	UID *string `json:"uid,omitempty"`
}

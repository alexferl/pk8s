// Code generated; DO NOT EDIT.

package k8s

// FlowSchemaSpecV1 FlowSchemaSpec describes how the FlowSchema's specification looks like.
type FlowSchemaSpecV1 struct {
	// FlowDistinguisherMethod specifies the method of a flow distinguisher.
	DistinguisherMethod *FlowDistinguisherMethodV1 `json:"distinguisherMethod,omitempty"`
	// `matchingPrecedence` is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.
	MatchingPrecedence *int32 `json:"matchingPrecedence,omitempty"`
	// PriorityLevelConfigurationReference contains information that points to the "request-priority" being used.
	PriorityLevelConfiguration PriorityLevelConfigurationReferenceV1 `json:"priorityLevelConfiguration"`
	// `rules` describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.
	Rules []PolicyRulesWithSubjectsV1 `json:"rules,omitempty"`
}

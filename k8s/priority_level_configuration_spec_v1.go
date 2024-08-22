// Code generated; DO NOT EDIT.

package k8s

// PriorityLevelConfigurationSpecV1 PriorityLevelConfigurationSpec specifies the configuration of a priority level.
type PriorityLevelConfigurationSpecV1 struct {
	// ExemptPriorityLevelConfiguration describes the configurable aspects of the handling of exempt requests. In the mandatory exempt configuration object the values in the fields here can be modified by authorized users, unlike the rest of the `spec`.
	Exempt *ExemptPriorityLevelConfigurationV1 `json:"exempt,omitempty"`
	// LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:   - How are requests for this priority level limited?   - What should be done with requests that exceed the limit?
	Limited *LimitedPriorityLevelConfigurationV1 `json:"limited,omitempty"`
	// `type` indicates whether this priority level is subject to limitation on request execution.  A value of `"Exempt"` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `"Limited"` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.
	Type string `json:"type"`
}

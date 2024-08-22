// Code generated; DO NOT EDIT.

package k8s

// SubjectAccessReviewSpecV1 SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
type SubjectAccessReviewSpecV1 struct {
	// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
	Extra *map[string][]string `json:"extra,omitempty"`
	// Groups is the groups you're testing for.
	Groups []string `json:"groups,omitempty"`
	// NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
	NonResourceAttributes *NonResourceAttributesV1 `json:"nonResourceAttributes,omitempty"`
	// ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
	ResourceAttributes *ResourceAttributesV1 `json:"resourceAttributes,omitempty"`
	// UID information about the requesting user.
	UID *string `json:"uid,omitempty"`
	// User is the user you're testing for. If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups
	User *string `json:"user,omitempty"`
}

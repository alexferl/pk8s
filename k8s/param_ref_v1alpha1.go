// Code generated; DO NOT EDIT.

package k8s

// ParamRefV1alpha1 ParamRef describes how to locate the params to be used as input to expressions of rules applied by a policy binding.
type ParamRefV1alpha1 struct {
	// `name` is the name of the resource being referenced.  `name` and `selector` are mutually exclusive properties. If one is set, the other must be unset.
	Name *string `json:"name,omitempty"`
	// namespace is the namespace of the referenced resource. Allows limiting the search for params to a specific namespace. Applies to both `name` and `selector` fields.  A per-namespace parameter may be used by specifying a namespace-scoped `paramKind` in the policy and leaving this field empty.  - If `paramKind` is cluster-scoped, this field MUST be unset. Setting this field results in a configuration error.  - If `paramKind` is namespace-scoped, the namespace of the object being evaluated for admission will be used when this field is left unset. Take care that if this is left empty the binding must not match any cluster-scoped resources, which will result in an error.
	Namespace *string `json:"namespace,omitempty"`
	// `parameterNotFoundAction` controls the behavior of the binding when the resource exists, and name or selector is valid, but there are no parameters matched by the binding. If the value is set to `Allow`, then no matched parameters will be treated as successful validation by the binding. If set to `Deny`, then no matched parameters will be subject to the `failurePolicy` of the policy.  Allowed values are `Allow` or `Deny` Default to `Deny`
	ParameterNotFoundAction *string `json:"parameterNotFoundAction,omitempty"`
	// A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
	Selector *LabelSelectorV1 `json:"selector,omitempty"`
}

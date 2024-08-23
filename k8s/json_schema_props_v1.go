// Code generated by pk8s; DO NOT EDIT.

package k8s

// JSONSchemaPropsV1 JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
type JSONSchemaPropsV1 struct {
	Ref    *string `json:"$ref,omitempty"`
	Schema *string `json:"$schema,omitempty"`
	// JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property.
	AdditionalItems map[string]interface{} `json:"additionalItems,omitempty"`
	// JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property.
	AdditionalProperties map[string]interface{} `json:"additionalProperties,omitempty"`
	AllOf                []JSONSchemaPropsV1    `json:"allOf,omitempty"`
	AnyOf                []JSONSchemaPropsV1    `json:"anyOf,omitempty"`
	// JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.
	Default      map[string]interface{}       `json:"default,omitempty"`
	Definitions  map[string]JSONSchemaPropsV1 `json:"definitions,omitempty"`
	Dependencies *map[string]interface{}      `json:"dependencies,omitempty"`
	Description  *string                      `json:"description,omitempty"`
	Enum         []map[string]interface{}     `json:"enum,omitempty"`
	// JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.
	Example          map[string]interface{} `json:"example,omitempty"`
	ExclusiveMaximum *bool                  `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum *bool                  `json:"exclusiveMinimum,omitempty"`
	// ExternalDocumentation allows referencing an external resource for extended documentation.
	ExternalDocs *ExternalDocumentationV1 `json:"externalDocs,omitempty"`
	// format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like "0321751043" or "978-0321751041" - isbn10: an ISBN10 number string like "0321751043" - isbn13: an ISBN13 number string like "978-0321751041" - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^\d{3}[- ]?\d{2}[- ]?\d{4}$ - hexcolor: an hexadecimal color code like "#FFFFFF: following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like "rgb(255,255,2559" - byte: base64 encoded binary data - password: any kind of string - date: a date string like "2006-01-02" as defined by full-date in RFC3339 - duration: a duration string like "22 ns" as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like "2014-12-15T19:30:20.000Z" as defined by date-time in RFC3339.
	Format *string `json:"format,omitempty"`
	ID     *string `json:"id,omitempty"`
	// JSONSchemaPropsOrArray represents a value that can either be a JSONSchemaProps or an array of JSONSchemaProps. Mainly here for serialization purposes.
	Items         map[string]interface{} `json:"items,omitempty"`
	MaxItems      *int64                 `json:"maxItems,omitempty"`
	MaxLength     *int64                 `json:"maxLength,omitempty"`
	MaxProperties *int64                 `json:"maxProperties,omitempty"`
	Maximum       *float64               `json:"maximum,omitempty"`
	MinItems      *int64                 `json:"minItems,omitempty"`
	MinLength     *int64                 `json:"minLength,omitempty"`
	MinProperties *int64                 `json:"minProperties,omitempty"`
	Minimum       *float64               `json:"minimum,omitempty"`
	MultipleOf    *float64               `json:"multipleOf,omitempty"`
	// JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
	Not               *JSONSchemaPropsV1           `json:"not,omitempty"`
	Nullable          *bool                        `json:"nullable,omitempty"`
	OneOf             []JSONSchemaPropsV1          `json:"oneOf,omitempty"`
	Pattern           *string                      `json:"pattern,omitempty"`
	PatternProperties map[string]JSONSchemaPropsV1 `json:"patternProperties,omitempty"`
	Properties        map[string]JSONSchemaPropsV1 `json:"properties,omitempty"`
	Required          []string                     `json:"required,omitempty"`
	Title             *string                      `json:"title,omitempty"`
	Type              *string                      `json:"type,omitempty"`
	UniqueItems       *bool                        `json:"uniqueItems,omitempty"`
	// x-kubernetes-embedded-resource defines that the value is an embedded Kubernetes runtime.Object, with TypeMeta and ObjectMeta. The type must be object. It is allowed to further restrict the embedded object. kind, apiVersion and metadata are validated automatically. x-kubernetes-preserve-unknown-fields is allowed to be true, but does not have to be if the object is fully specified (up to kind, apiVersion, metadata).
	XKubernetesEmbeddedResource *bool `json:"x-kubernetes-embedded-resource,omitempty"`
	// x-kubernetes-int-or-string specifies that this value is either an integer or a string. If this is true, an empty type is allowed and type as child of anyOf is permitted if following one of the following patterns:  1) anyOf:    - type: integer    - type: string 2) allOf:    - anyOf:      - type: integer      - type: string    - ... zero or more
	XKubernetesIntOrString *bool `json:"x-kubernetes-int-or-string,omitempty"`
	// x-kubernetes-list-map-keys annotates an array with the x-kubernetes-list-type `map` by specifying the keys used as the index of the map.  This tag MUST only be used on lists that have the "x-kubernetes-list-type" extension set to "map". Also, the values specified for this attribute must be a scalar typed field of the child structure (no nesting is supported).  The properties specified must either be required or have a default value, to ensure those properties are present for all list items.
	XKubernetesListMapKeys []string `json:"x-kubernetes-list-map-keys,omitempty"`
	// x-kubernetes-list-type annotates an array to further describe its topology. This extension must only be used on lists and may have 3 possible values:  1) `atomic`: the list is treated as a single entity, like a scalar.      Atomic lists will be entirely replaced when updated. This extension      may be used on any type of list (struct, scalar, ...). 2) `set`:      Sets are lists that must not have multiple items with the same value. Each      value must be a scalar, an object with x-kubernetes-map-type `atomic` or an      array with x-kubernetes-list-type `atomic`. 3) `map`:      These lists are like maps in that their elements have a non-index key      used to identify them. Order is preserved upon merge. The map tag      must only be used on a list with elements of type object. Defaults to atomic for arrays.
	XKubernetesListType *string `json:"x-kubernetes-list-type,omitempty"`
	// x-kubernetes-map-type annotates an object to further describe its topology. This extension must only be used when type is object and may have 2 possible values:  1) `granular`:      These maps are actual maps (key-value pairs) and each fields are independent      from each other (they can each be manipulated by separate actors). This is      the default behaviour for all maps. 2) `atomic`: the list is treated as a single entity, like a scalar.      Atomic maps will be entirely replaced when updated.
	XKubernetesMapType *string `json:"x-kubernetes-map-type,omitempty"`
	// x-kubernetes-preserve-unknown-fields stops the API server decoding step from pruning fields which are not specified in the validation schema. This affects fields recursively, but switches back to normal pruning behaviour if nested properties or additionalProperties are specified in the schema. This can either be true or undefined. False is forbidden.
	XKubernetesPreserveUnknownFields *bool `json:"x-kubernetes-preserve-unknown-fields,omitempty"`
	// x-kubernetes-validations describes a list of validation rules written in the CEL expression language.
	XKubernetesValidations []ValidationRuleV1 `json:"x-kubernetes-validations,omitempty"`
}

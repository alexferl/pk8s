// Code generated; DO NOT EDIT.

package k8s

// CustomResourceDefinitionVersionV1 CustomResourceDefinitionVersion describes a version for CRD.
type CustomResourceDefinitionVersionV1 struct {
	// additionalPrinterColumns specifies additional columns returned in Table output. See https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details. If no columns are specified, a single column displaying the age of the custom resource is used.
	AdditionalPrinterColumns []CustomResourceColumnDefinitionV1 `json:"additionalPrinterColumns,omitempty"`
	// deprecated indicates this version of the custom resource API is deprecated. When set to true, API requests to this version receive a warning header in the server response. Defaults to false.
	Deprecated *bool `json:"deprecated,omitempty"`
	// deprecationWarning overrides the default warning returned to API clients. May only be set when `deprecated` is true. The default warning indicates this version is deprecated and recommends use of the newest served version of equal or greater stability, if one exists.
	DeprecationWarning *string `json:"deprecationWarning,omitempty"`
	// name is the version name, e.g. “v1”, “v2beta1”, etc. The custom resources are served under this version at `/apis/<group>/<version>/...` if `served` is true.
	Name string `json:"name"`
	// CustomResourceValidation is a list of validation methods for CustomResources.
	Schema *CustomResourceValidationV1 `json:"schema,omitempty"`
	// selectableFields specifies paths to fields that may be used as field selectors. A maximum of 8 selectable fields are allowed. See https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors
	SelectableFields []SelectableFieldV1 `json:"selectableFields,omitempty"`
	// served is a flag enabling/disabling this version from being served via REST APIs
	Served bool `json:"served"`
	// storage indicates this version should be used when persisting custom resources to storage. There must be exactly one version with storage=true.
	Storage bool `json:"storage"`
	// CustomResourceSubresources defines the status and scale subresources for CustomResources.
	Subresources *CustomResourceSubresourcesV1 `json:"subresources,omitempty"`
}

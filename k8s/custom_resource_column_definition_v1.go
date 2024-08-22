// Code generated; DO NOT EDIT.

package k8s

// CustomResourceColumnDefinitionV1 CustomResourceColumnDefinition specifies a column for server side printing.
type CustomResourceColumnDefinitionV1 struct {
	// description is a human readable description of this column.
	Description *string `json:"description,omitempty"`
	// format is an optional OpenAPI type definition for this column. The 'name' format is applied to the primary identifier column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.
	Format *string `json:"format,omitempty"`
	// jsonPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the value for this column.
	JSONPath string `json:"jsonPath"`
	// name is a human readable name for the column.
	Name string `json:"name"`
	// priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0.
	Priority *int32 `json:"priority,omitempty"`
	// type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.
	Type string `json:"type"`
}

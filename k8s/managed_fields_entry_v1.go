// Code generated by pk8s; DO NOT EDIT.

package k8s

import (
	"time"
)

// ManagedFieldsEntryV1 ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
type ManagedFieldsEntryV1 struct {
	// APIVersion defines the version of this resource that this field set applies to. The format is "group/version" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
	APIVersion *string `json:"apiVersion,omitempty"`
	// FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: "FieldsV1"
	FieldsType *string `json:"fieldsType,omitempty"`
	// FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format.  Each key is either a '.' representing the field itself, and will always map to an empty set, or a string representing a sub-field or item. The string will follow one of these four formats: 'f:<name>', where <name> is the name of a field in a struct, or key in a map 'v:<value>', where <value> is the exact json formatted value of a list item 'i:<index>', where <index> is position of a item in a list 'k:<keys>', where <keys> is a map of  a list item's key fields to their unique values If a key maps to an empty Fields value, the field that key represents is part of the set.  The exact format is defined in sigs.k8s.io/structured-merge-diff
	FieldsV1 map[string]interface{} `json:"fieldsV1,omitempty"`
	// Manager is an identifier of the workflow managing these fields.
	Manager *string `json:"manager,omitempty"`
	// Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.
	Operation *string `json:"operation,omitempty"`
	// Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource.
	Subresource *string `json:"subresource,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	Time *time.Time `json:"time,omitempty"`
}

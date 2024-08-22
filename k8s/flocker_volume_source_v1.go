// Code generated; DO NOT EDIT.

package k8s

// FlockerVolumeSourceV1 Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
type FlockerVolumeSourceV1 struct {
	// datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
	DatasetName *string `json:"datasetName,omitempty"`
	// datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset
	DatasetUUID *string `json:"datasetUUID,omitempty"`
}

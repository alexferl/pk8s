// Code generated; DO NOT EDIT.

package k8s

// AzureFileVolumeSourceV1 AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type AzureFileVolumeSourceV1 struct {
	// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// secretName is the  name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName"`
	// shareName is the azure share Name
	ShareName string `json:"shareName"`
}

// Code generated; DO NOT EDIT.

package k8s

// AzureFilePersistentVolumeSourceV1 AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type AzureFilePersistentVolumeSourceV1 struct {
	// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// secretName is the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName"`
	// secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
	SecretNamespace *string `json:"secretNamespace,omitempty"`
	// shareName is the azure Share Name
	ShareName string `json:"shareName"`
}

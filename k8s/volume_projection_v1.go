// Code generated; DO NOT EDIT.

package k8s

// VolumeProjectionV1 Projection that may be projected along with other supported volume types. Exactly one of these fields must be set.
type VolumeProjectionV1 struct {
	// ClusterTrustBundleProjection describes how to select a set of ClusterTrustBundle objects and project their contents into the pod filesystem.
	ClusterTrustBundle *ClusterTrustBundleProjectionV1 `json:"clusterTrustBundle,omitempty"`
	// Adapts a ConfigMap into a projected volume.  The contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.
	ConfigMap *ConfigMapProjectionV1 `json:"configMap,omitempty"`
	// Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume source without the default mode.
	DownwardAPI *DownwardAPIProjectionV1 `json:"downwardAPI,omitempty"`
	// Adapts a secret into a projected volume.  The contents of the target Secret's Data field will be presented in a projected volume as files using the keys in the Data field as the file names. Note that this is identical to a secret volume source without the default mode.
	Secret *SecretProjectionV1 `json:"secret,omitempty"`
	// ServiceAccountTokenProjection represents a projected service account token volume. This projection can be used to insert a service account token into the pods runtime filesystem for use against APIs (Kubernetes API Server or otherwise).
	ServiceAccountToken *ServiceAccountTokenProjectionV1 `json:"serviceAccountToken,omitempty"`
}

// Code generated; DO NOT EDIT.

package k8s

// DownwardAPIVolumeFileV1 DownwardAPIVolumeFile represents information to create the file containing the pod field
type DownwardAPIVolumeFileV1 struct {
	// ObjectFieldSelector selects an APIVersioned field of an object.
	FieldRef *ObjectFieldSelectorV1 `json:"fieldRef,omitempty"`
	// Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Mode *int32 `json:"mode,omitempty"`
	// Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	Path string `json:"path"`
	// ResourceFieldSelector represents container resources (cpu, memory) and their output format
	ResourceFieldRef *ResourceFieldSelectorV1 `json:"resourceFieldRef,omitempty"`
}

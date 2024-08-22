// Code generated; DO NOT EDIT.

package k8s

// DownwardAPIVolumeSourceV1 DownwardAPIVolumeSource represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.
type DownwardAPIVolumeSourceV1 struct {
	// Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode *int32 `json:"defaultMode,omitempty"`
	// Items is a list of downward API volume file
	Items []DownwardAPIVolumeFileV1 `json:"items,omitempty"`
}

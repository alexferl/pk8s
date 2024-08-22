// Code generated; DO NOT EDIT.

package k8s

// DownwardAPIProjectionV1 Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume source without the default mode.
type DownwardAPIProjectionV1 struct {
	// Items is a list of DownwardAPIVolume file
	Items []DownwardAPIVolumeFileV1 `json:"items,omitempty"`
}

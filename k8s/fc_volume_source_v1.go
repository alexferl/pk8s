// Code generated; DO NOT EDIT.

package k8s

// FCVolumeSourceV1 Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
type FCVolumeSourceV1 struct {
	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// lun is Optional: FC target lun number
	Lun *int32 `json:"lun,omitempty"`
	// readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// targetWWNs is Optional: FC target worldwide names (WWNs)
	TargetWwNs []string `json:"targetWWNs,omitempty"`
	// wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	Wwids []string `json:"wwids,omitempty"`
}

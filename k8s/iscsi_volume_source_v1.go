// Code generated; DO NOT EDIT.

package k8s

// ISCSIVolumeSourceV1 Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
type ISCSIVolumeSourceV1 struct {
	// chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
	ChapAuthDiscovery *bool `json:"chapAuthDiscovery,omitempty"`
	// chapAuthSession defines whether support iSCSI Session CHAP authentication
	ChapAuthSession *bool `json:"chapAuthSession,omitempty"`
	// fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	FsType *string `json:"fsType,omitempty"`
	// initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	InitiatorName *string `json:"initiatorName,omitempty"`
	// iqn is the target iSCSI Qualified Name.
	Iqn string `json:"iqn"`
	// iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
	IscsiInterface *string `json:"iscsiInterface,omitempty"`
	// lun represents iSCSI Target Lun number.
	Lun int32 `json:"lun"`
	// portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Portals []string `json:"portals,omitempty"`
	// readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
	SecretRef *LocalObjectReferenceV1 `json:"secretRef,omitempty"`
	// targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	TargetPortal string `json:"targetPortal"`
}

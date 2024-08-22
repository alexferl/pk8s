// Code generated; DO NOT EDIT.

package k8s

// CSINodeSpecV1 CSINodeSpec holds information about the specification of all CSI drivers installed on a node
type CSINodeSpecV1 struct {
	// drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this can become empty.
	Drivers []CSINodeDriverV1 `json:"drivers"`
}

// Code generated; DO NOT EDIT.

package k8s

// ContainerResizePolicyV1 ContainerResizePolicy represents resource resize policy for the container.
type ContainerResizePolicyV1 struct {
	// Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.
	ResourceName string `json:"resourceName"`
	// Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.
	RestartPolicy string `json:"restartPolicy"`
}

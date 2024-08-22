// Code generated; DO NOT EDIT.

package k8s

// ContainerStatusV1 ContainerStatus contains details for the current status of this container.
type ContainerStatusV1 struct {
	// AllocatedResources represents the compute resources allocated for this container by the node. Kubelet sets this value to Container.Resources.Requests upon successful pod admission and after successfully admitting desired pod resize.
	AllocatedResources *map[string]string `json:"allocatedResources,omitempty"`
	// AllocatedResourcesStatus represents the status of various resources allocated for this Pod.
	AllocatedResourcesStatus []ResourceStatusV1 `json:"allocatedResourcesStatus,omitempty"`
	// ContainerID is the ID of the container in the format '<type>://<container_id>'. Where type is a container runtime identifier, returned from Version call of CRI API (for example "containerd").
	ContainerID *string `json:"containerID,omitempty"`
	// Image is the name of container image that the container is running. The container image may not match the image used in the PodSpec, as it may have been resolved by the runtime. More info: https://kubernetes.io/docs/concepts/containers/images.
	Image string `json:"image"`
	// ImageID is the image ID of the container's image. The image ID may not match the image ID of the image used in the PodSpec, as it may have been resolved by the runtime.
	ImageID string `json:"imageID"`
	// ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
	LastState *ContainerStateV1 `json:"lastState,omitempty"`
	// Name is a DNS_LABEL representing the unique name of the container. Each container in a pod must have a unique name across all container types. Cannot be updated.
	Name string `json:"name"`
	// Ready specifies whether the container is currently passing its readiness check. The value will change as readiness probes keep executing. If no readiness probes are specified, this field defaults to true once the container is fully started (see Started field).  The value is typically used to determine whether a container is ready to accept traffic.
	Ready bool `json:"ready"`
	// ResourceRequirements describes the compute resource requirements.
	Resources *ResourceRequirementsV1 `json:"resources,omitempty"`
	// RestartCount holds the number of times the container has been restarted. Kubelet makes an effort to always increment the value, but there are cases when the state may be lost due to node restarts and then the value may be reset to 0. The value is never negative.
	RestartCount int32 `json:"restartCount"`
	// Started indicates whether the container has finished its postStart lifecycle hook and passed its startup probe. Initialized as false, becomes true after startupProbe is considered successful. Resets to false when the container is restarted, or if kubelet loses state temporarily. In both cases, startup probes will run again. Is always true when no startupProbe is defined and container is running and has passed the postStart lifecycle hook. The null value must be treated the same as false.
	Started *bool `json:"started,omitempty"`
	// ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
	State *ContainerStateV1 `json:"state,omitempty"`
	// ContainerUser represents user identity information
	User *ContainerUserV1 `json:"user,omitempty"`
	// Status of volume mounts.
	VolumeMounts []VolumeMountStatusV1 `json:"volumeMounts,omitempty"`
}

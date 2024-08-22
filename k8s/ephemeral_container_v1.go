// Code generated; DO NOT EDIT.

package k8s

// EphemeralContainerV1 An EphemeralContainer is a temporary container that you may add to an existing Pod for user-initiated activities such as debugging. Ephemeral containers have no resource or scheduling guarantees, and they will not be restarted when they exit or when a Pod is removed or restarted. The kubelet may evict a Pod if an ephemeral container causes the Pod to exceed its resource allocation.  To add an ephemeral container, use the ephemeralcontainers subresource of an existing Pod. Ephemeral containers may not be removed or restarted.
type EphemeralContainerV1 struct {
	// Arguments to the entrypoint. The image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Args []string `json:"args,omitempty"`
	// Entrypoint array. Not executed within a shell. The image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Command []string `json:"command,omitempty"`
	// List of environment variables to set in the container. Cannot be updated.
	Env []EnvVarV1 `json:"env,omitempty"`
	// List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	EnvFrom []EnvFromSourceV1 `json:"envFrom,omitempty"`
	// Container image name. More info: https://kubernetes.io/docs/concepts/containers/images
	Image *string `json:"image,omitempty"`
	// Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	// Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.
	Lifecycle *LifecycleV1 `json:"lifecycle,omitempty"`
	// Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	LivenessProbe *ProbeV1 `json:"livenessProbe,omitempty"`
	// Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.
	Name string `json:"name"`
	// Ports are not allowed for ephemeral containers.
	Ports []ContainerPortV1 `json:"ports,omitempty"`
	// Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	ReadinessProbe *ProbeV1 `json:"readinessProbe,omitempty"`
	// Resources resize policy for the container.
	ResizePolicy []ContainerResizePolicyV1 `json:"resizePolicy,omitempty"`
	// ResourceRequirements describes the compute resource requirements.
	Resources *ResourceRequirementsV1 `json:"resources,omitempty"`
	// Restart policy for the container to manage the restart behavior of each container within a pod. This may only be set for init containers. You cannot set this field on ephemeral containers.
	RestartPolicy *string `json:"restartPolicy,omitempty"`
	// SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.
	SecurityContext *SecurityContextV1 `json:"securityContext,omitempty"`
	// Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
	StartupProbe *ProbeV1 `json:"startupProbe,omitempty"`
	// Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
	Stdin *bool `json:"stdin,omitempty"`
	// Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
	StdinOnce *bool `json:"stdinOnce,omitempty"`
	// If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container uses the namespaces configured in the Pod spec.  The container runtime must implement support for this feature. If the runtime does not support namespace targeting then the result of setting this field is undefined.
	TargetContainerName *string `json:"targetContainerName,omitempty"`
	// Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
	TerminationMessagePath *string `json:"terminationMessagePath,omitempty"`
	// Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	TerminationMessagePolicy *string `json:"terminationMessagePolicy,omitempty"`
	// Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	Tty *bool `json:"tty,omitempty"`
	// volumeDevices is the list of block devices to be used by the container.
	VolumeDevices []VolumeDeviceV1 `json:"volumeDevices,omitempty"`
	// Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. Cannot be updated.
	VolumeMounts []VolumeMountV1 `json:"volumeMounts,omitempty"`
	// Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	WorkingDir *string `json:"workingDir,omitempty"`
}

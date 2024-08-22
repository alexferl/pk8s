// Code generated; DO NOT EDIT.

package k8s

// SecurityContextV1 SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.
type SecurityContextV1 struct {
	// AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.
	AllowPrivilegeEscalation *bool `json:"allowPrivilegeEscalation,omitempty"`
	// AppArmorProfile defines a pod or container's AppArmor settings.
	AppArmorProfile *AppArmorProfileV1 `json:"appArmorProfile,omitempty"`
	// Adds and removes POSIX capabilities from running containers.
	Capabilities *CapabilitiesV1 `json:"capabilities,omitempty"`
	// Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.
	Privileged *bool `json:"privileged,omitempty"`
	// procMount denotes the type of proc mount to use for the containers. The default value is Default which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.
	ProcMount *string `json:"procMount,omitempty"`
	// Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.
	ReadOnlyRootFilesystem *bool `json:"readOnlyRootFilesystem,omitempty"`
	// The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.
	RunAsGroup *int64 `json:"runAsGroup,omitempty"`
	// Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty"`
	// The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.
	RunAsUser *int64 `json:"runAsUser,omitempty"`
	// SELinuxOptions are the labels to be applied to the container
	SeLinuxOptions *SELinuxOptionsV1 `json:"seLinuxOptions,omitempty"`
	// SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.
	SeccompProfile *SeccompProfileV1 `json:"seccompProfile,omitempty"`
	// WindowsSecurityContextOptions contain Windows-specific options and credentials.
	WindowsOptions *WindowsSecurityContextOptionsV1 `json:"windowsOptions,omitempty"`
}

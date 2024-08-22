// Code generated; DO NOT EDIT.

package k8s

// PodSecurityContextV1 PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.
type PodSecurityContextV1 struct {
	// AppArmorProfile defines a pod or container's AppArmor settings.
	AppArmorProfile *AppArmorProfileV1 `json:"appArmorProfile,omitempty"`
	// A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.
	FsGroup *int64 `json:"fsGroup,omitempty"`
	// fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are "OnRootMismatch" and "Always". If not specified, "Always" is used. Note that this field cannot be set when spec.os.name is windows.
	FsGroupChangePolicy *string `json:"fsGroupChangePolicy,omitempty"`
	// The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.
	RunAsGroup *int64 `json:"runAsGroup,omitempty"`
	// Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty"`
	// The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.
	RunAsUser *int64 `json:"runAsUser,omitempty"`
	// SELinuxOptions are the labels to be applied to the container
	SeLinuxOptions *SELinuxOptionsV1 `json:"seLinuxOptions,omitempty"`
	// SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.
	SeccompProfile *SeccompProfileV1 `json:"seccompProfile,omitempty"`
	// A list of groups applied to the first process run in each container, in addition to the container's primary GID and fsGroup (if specified).  If the SupplementalGroupsPolicy feature is enabled, the supplementalGroupsPolicy field determines whether these are in addition to or instead of any group memberships defined in the container image. If unspecified, no additional groups are added, though group memberships defined in the container image may still be used, depending on the supplementalGroupsPolicy field. Note that this field cannot be set when spec.os.name is windows.
	SupplementalGroups []int64 `json:"supplementalGroups,omitempty"`
	// Defines how supplemental groups of the first container processes are calculated. Valid values are "Merge" and "Strict". If not specified, "Merge" is used. (Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled and the container runtime must implement support for this feature. Note that this field cannot be set when spec.os.name is windows.
	SupplementalGroupsPolicy *string `json:"supplementalGroupsPolicy,omitempty"`
	// Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.
	Sysctls []SysctlV1 `json:"sysctls,omitempty"`
	// WindowsSecurityContextOptions contain Windows-specific options and credentials.
	WindowsOptions *WindowsSecurityContextOptionsV1 `json:"windowsOptions,omitempty"`
}

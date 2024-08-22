// Code generated; DO NOT EDIT.

package k8s

// PodSpecV1 PodSpec is a description of a pod.
type PodSpecV1 struct {
	// Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty"`
	// Affinity is a group of affinity scheduling rules.
	Affinity *AffinityV1 `json:"affinity,omitempty"`
	// AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
	AutomountServiceAccountToken *bool `json:"automountServiceAccountToken,omitempty"`
	// List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.
	Containers []ContainerV1 `json:"containers"`
	// PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.
	DNSConfig *PodDNSConfigV1 `json:"dnsConfig,omitempty"`
	// Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
	DNSPolicy *string `json:"dnsPolicy,omitempty"`
	// EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.
	EnableServiceLinks *bool `json:"enableServiceLinks,omitempty"`
	// List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing pod to perform user-initiated actions such as debugging. This list cannot be specified when creating a pod, and it cannot be modified by updating the pod spec. In order to add an ephemeral container to an existing pod, use the pod's ephemeralcontainers subresource.
	EphemeralContainers []EphemeralContainerV1 `json:"ephemeralContainers,omitempty"`
	// HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified.
	HostAliases []HostAliasV1 `json:"hostAliases,omitempty"`
	// Use the host's ipc namespace. Optional: Default to false.
	HostIpc *bool `json:"hostIPC,omitempty"`
	// Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.
	HostNetwork *bool `json:"hostNetwork,omitempty"`
	// Use the host's pid namespace. Optional: Default to false.
	HostPid *bool `json:"hostPID,omitempty"`
	// Use the host's user namespace. Optional: Default to true. If set to true or not present, the pod will be run in the host user namespace, useful for when the pod needs a feature only available to the host user namespace, such as loading a kernel module with CAP_SYS_MODULE. When set to false, a new userns is created for the pod. Setting false is useful for mitigating container breakout vulnerabilities even allowing users to run their containers as root without actually having root privileges on the host. This field is alpha-level and is only honored by servers that enable the UserNamespacesSupport feature.
	HostUsers *bool `json:"hostUsers,omitempty"`
	// Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.
	Hostname *string `json:"hostname,omitempty"`
	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets []LocalObjectReferenceV1 `json:"imagePullSecrets,omitempty"`
	// List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	InitContainers []ContainerV1 `json:"initContainers,omitempty"`
	// NodeName indicates in which node this pod is scheduled. If empty, this pod is a candidate for scheduling by the scheduler defined in schedulerName. Once this field is set, the kubelet for this node becomes responsible for the lifecycle of this pod. This field should not be used to express a desire for the pod to be scheduled on a specific node. https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodename
	NodeName *string `json:"nodeName,omitempty"`
	// NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector *map[string]string `json:"nodeSelector,omitempty"`
	// PodOS defines the OS parameters of a pod.
	Os *PodOSV1 `json:"os,omitempty"`
	// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md
	Overhead *map[string]string `json:"overhead,omitempty"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
	PreemptionPolicy *string `json:"preemptionPolicy,omitempty"`
	// The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.
	Priority *int32 `json:"priority,omitempty"`
	// If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
	PriorityClassName *string `json:"priorityClassName,omitempty"`
	// If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates
	ReadinessGates []PodReadinessGateV1 `json:"readinessGates,omitempty"`
	// ResourceClaims defines which ResourceClaims must be allocated and reserved before the Pod is allowed to start. The resources will be made available to those containers which consume them by name.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable.
	ResourceClaims []PodResourceClaimV1 `json:"resourceClaims,omitempty"`
	// Restart policy for all containers within the pod. One of Always, OnFailure, Never. In some contexts, only a subset of those values may be permitted. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	RestartPolicy *string `json:"restartPolicy,omitempty"`
	// RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class
	RuntimeClassName *string `json:"runtimeClassName,omitempty"`
	// If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.
	SchedulerName *string `json:"schedulerName,omitempty"`
	// SchedulingGates is an opaque list of values that if specified will block scheduling the pod. If schedulingGates is not empty, the pod will stay in the SchedulingGated state and the scheduler will not attempt to schedule the pod.  SchedulingGates can only be set at pod creation time, and be removed only afterwards.
	SchedulingGates []PodSchedulingGateV1 `json:"schedulingGates,omitempty"`
	// PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.
	SecurityContext *PodSecurityContextV1 `json:"securityContext,omitempty"`
	// DeprecatedServiceAccount is a deprecated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.
	ServiceAccount *string `json:"serviceAccount,omitempty"`
	// ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	ServiceAccountName *string `json:"serviceAccountName,omitempty"`
	// If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname). In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters to FQDN. If a pod does not have FQDN, this has no effect. Default to false.
	SetHostnameAsFqdn *bool `json:"setHostnameAsFQDN,omitempty"`
	// Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.
	ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty"`
	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all.
	Subdomain *string `json:"subdomain,omitempty"`
	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`
	// If specified, the pod's tolerations.
	Tolerations []TolerationV1 `json:"tolerations,omitempty"`
	// TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.
	TopologySpreadConstraints []TopologySpreadConstraintV1 `json:"topologySpreadConstraints,omitempty"`
	// List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes
	Volumes []VolumeV1 `json:"volumes,omitempty"`
}

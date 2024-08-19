package k8s

func SetApiVersionAndKind(i any) {
	switch v := i.(type) {
	case *APIGroupV1:
		v.ApiVersion = String("v1")
		v.Kind = String("APIGroup")
	case *APIServiceV1:
		v.ApiVersion = String("apiregistration.k8s.io/v1")
		v.Kind = String("APIService")
	case *APIVersionsV1:
		v.ApiVersion = String("v1")
		v.Kind = String("APIVersions")
	case *BindingV1:
		v.ApiVersion = String("v1")
		v.Kind = String("Binding")
	case *CSIDriverV1:
		v.ApiVersion = String("storage.k8s.io/v1")
		v.Kind = String("CSIDriver")
	case *CSINodeV1:
		v.ApiVersion = String("storage.k8s.io/v1")
		v.Kind = String("CSINode")
	case *CSIStorageCapacityV1:
		v.ApiVersion = String("storage.k8s.io/v1")
		v.Kind = String("CSIStorageCapacity")
	case *CertificateSigningRequestV1:
		v.ApiVersion = String("certificates.k8s.io/v1")
		v.Kind = String("CertificateSigningRequest")
	case *ClusterRoleBindingV1:
		v.ApiVersion = String("rbac.authorization.k8s.io/v1")
		v.Kind = String("ClusterRoleBinding")
	case *ClusterRoleV1:
		v.ApiVersion = String("rbac.authorization.k8s.io/v1")
		v.Kind = String("ClusterRole")
	case *ClusterTrustBundleV1alpha1:
		v.ApiVersion = String("certificates.k8s.io/v1alpha1")
		v.Kind = String("ClusterTrustBundle")
	case *ComponentStatusV1:
		v.ApiVersion = String("v1")
		v.Kind = String("ComponentStatus")
	case *ConfigMapV1:
		v.ApiVersion = String("v1")
		v.Kind = String("ConfigMap")
	case *ControllerRevisionV1:
		v.ApiVersion = String("apps/v1")
		v.Kind = String("ControllerRevision")
	case *CronJobV1:
		v.ApiVersion = String("batch/v1")
		v.Kind = String("CronJob")
	case *CustomResourceDefinitionV1:
		v.ApiVersion = String("apiextensions.k8s.io/v1")
		v.Kind = String("CustomResourceDefinition")
	case *DaemonSetV1:
		v.ApiVersion = String("apps/v1")
		v.Kind = String("DaemonSet")
	case *DeploymentV1:
		v.ApiVersion = String("apps/v1")
		v.Kind = String("Deployment")
	case *EndpointSliceV1:
		v.ApiVersion = String("discovery.k8s.io/v1")
		v.Kind = String("EndpointSlice")
	case *EndpointsV1:
		v.ApiVersion = String("v1")
		v.Kind = String("Endpoints")
	case *EventV1:
		v.ApiVersion = String("events.k8s.io/v1")
		v.Kind = String("Event")
	case *EvictionV1:
		v.ApiVersion = String("policy/v1")
		v.Kind = String("Eviction")
	case *FlowSchemaV1:
		v.ApiVersion = String("flowcontrol.apiserver.k8s.io/v1")
		v.Kind = String("FlowSchema")
	case *FlowSchemaV1beta3:
		v.ApiVersion = String("flowcontrol.apiserver.k8s.io/v1beta3")
		v.Kind = String("FlowSchema")
	case *HorizontalPodAutoscalerV1:
		v.ApiVersion = String("autoscaling/v1")
		v.Kind = String("HorizontalPodAutoscaler")
	case *HorizontalPodAutoscalerV2:
		v.ApiVersion = String("autoscaling/v2")
		v.Kind = String("HorizontalPodAutoscaler")
	case *IPAddressV1alpha1:
		v.ApiVersion = String("networking.k8s.io/v1alpha1")
		v.Kind = String("IPAddress")
	case *IngressClassV1:
		v.ApiVersion = String("networking.k8s.io/v1")
		v.Kind = String("IngressClass")
	case *IngressV1:
		v.ApiVersion = String("networking.k8s.io/v1")
		v.Kind = String("Ingress")
	case *JobV1:
		v.ApiVersion = String("batch/v1")
		v.Kind = String("Job")
	case *LeaseV1:
		v.ApiVersion = String("coordination.k8s.io/v1")
		v.Kind = String("Lease")
	case *LimitRangeV1:
		v.ApiVersion = String("v1")
		v.Kind = String("LimitRange")
	case *LocalSubjectAccessReviewV1:
		v.ApiVersion = String("authorization.k8s.io/v1")
		v.Kind = String("LocalSubjectAccessReview")
	case *MutatingWebhookConfigurationV1:
		v.ApiVersion = String("admissionregistration.k8s.io/v1")
		v.Kind = String("MutatingWebhookConfiguration")
	case *NamespaceV1:
		v.ApiVersion = String("v1")
		v.Kind = String("Namespace")
	case *NetworkPolicyV1:
		v.ApiVersion = String("networking.k8s.io/v1")
		v.Kind = String("NetworkPolicy")
	case *NodeV1:
		v.ApiVersion = String("v1")
		v.Kind = String("Node")
	case *PersistentVolumeClaimV1:
		v.ApiVersion = String("v1")
		v.Kind = String("PersistentVolumeClaim")
	case *PersistentVolumeV1:
		v.ApiVersion = String("v1")
		v.Kind = String("PersistentVolume")
	case *PodDisruptionBudgetV1:
		v.ApiVersion = String("policy/v1")
		v.Kind = String("PodDisruptionBudget")
	case *PodSchedulingContextV1alpha2:
		v.ApiVersion = String("resource.k8s.io/v1alpha2")
		v.Kind = String("PodSchedulingContext")
	case *PodTemplateV1:
		v.ApiVersion = String("v1")
		v.Kind = String("PodTemplate")
	case *PodV1:
		v.ApiVersion = String("v1")
		v.Kind = String("Pod")
	case *PriorityClassV1:
		v.ApiVersion = String("scheduling.k8s.io/v1")
		v.Kind = String("PriorityClass")
	case *PriorityLevelConfigurationV1:
		v.ApiVersion = String("flowcontrol.apiserver.k8s.io/v1")
		v.Kind = String("PriorityLevelConfiguration")
	case *PriorityLevelConfigurationV1beta3:
		v.ApiVersion = String("flowcontrol.apiserver.k8s.io/v1beta3")
		v.Kind = String("PriorityLevelConfiguration")
	case *ReplicaSetV1:
		v.ApiVersion = String("apps/v1")
		v.Kind = String("ReplicaSet")
	case *ReplicationControllerV1:
		v.ApiVersion = String("v1")
		v.Kind = String("ReplicationController")
	case *ResourceClaimParametersV1alpha2:
		v.ApiVersion = String("resource.k8s.io/v1alpha2")
		v.Kind = String("ResourceClaimParameters")
	case *ResourceClaimTemplateV1alpha2:
		v.ApiVersion = String("resource.k8s.io/v1alpha2")
		v.Kind = String("ResourceClaimTemplate")
	case *ResourceClaimV1alpha2:
		v.ApiVersion = String("resource.k8s.io/v1alpha2")
		v.Kind = String("ResourceClaim")
	case *ResourceClassParametersV1alpha2:
		v.ApiVersion = String("resource.k8s.io/v1alpha2")
		v.Kind = String("ResourceClassParameters")
	case *ResourceClassV1alpha2:
		v.ApiVersion = String("resource.k8s.io/v1alpha2")
		v.Kind = String("ResourceClass")
	case *ResourceQuotaV1:
		v.ApiVersion = String("v1")
		v.Kind = String("ResourceQuota")
	case *ResourceSliceV1alpha2:
		v.ApiVersion = String("resource.k8s.io/v1alpha2")
		v.Kind = String("ResourceSlice")
	case *RoleBindingV1:
		v.ApiVersion = String("rbac.authorization.k8s.io/v1")
		v.Kind = String("RoleBinding")
	case *RoleV1:
		v.ApiVersion = String("rbac.authorization.k8s.io/v1")
		v.Kind = String("Role")
	case *RuntimeClassV1:
		v.ApiVersion = String("node.k8s.io/v1")
		v.Kind = String("RuntimeClass")
	case *ScaleV1:
		v.ApiVersion = String("autoscaling/v1")
		v.Kind = String("Scale")
	case *SecretV1:
		v.ApiVersion = String("v1")
		v.Kind = String("Secret")
	case *SelfSubjectAccessReviewV1:
		v.ApiVersion = String("authorization.k8s.io/v1")
		v.Kind = String("SelfSubjectAccessReview")
	case *SelfSubjectReviewV1:
		v.ApiVersion = String("authentication.k8s.io/v1")
		v.Kind = String("SelfSubjectReview")
	case *SelfSubjectReviewV1alpha1:
		v.ApiVersion = String("authentication.k8s.io/v1alpha1")
		v.Kind = String("SelfSubjectReview")
	case *SelfSubjectReviewV1beta1:
		v.ApiVersion = String("authentication.k8s.io/v1beta1")
		v.Kind = String("SelfSubjectReview")
	case *SelfSubjectRulesReviewV1:
		v.ApiVersion = String("authorization.k8s.io/v1")
		v.Kind = String("SelfSubjectRulesReview")
	case *ServiceAccountV1:
		v.ApiVersion = String("v1")
		v.Kind = String("ServiceAccount")
	case *ServiceCIDRV1alpha1:
		v.ApiVersion = String("networking.k8s.io/v1alpha1")
		v.Kind = String("ServiceCIDR")
	case *ServiceV1:
		v.ApiVersion = String("v1")
		v.Kind = String("Service")
	case *StatefulSetV1:
		v.ApiVersion = String("apps/v1")
		v.Kind = String("StatefulSet")
	case *StatusV1:
		v.ApiVersion = String("v1")
		v.Kind = String("Status")
	case *StorageClassV1:
		v.ApiVersion = String("storage.k8s.io/v1")
		v.Kind = String("StorageClass")
	case *StorageVersionMigrationV1alpha1:
		v.ApiVersion = String("storagemigration.k8s.io/v1alpha1")
		v.Kind = String("StorageVersionMigration")
	case *StorageVersionV1alpha1:
		v.ApiVersion = String("internal.apiserver.k8s.io/v1alpha1")
		v.Kind = String("StorageVersion")
	case *SubjectAccessReviewV1:
		v.ApiVersion = String("authorization.k8s.io/v1")
		v.Kind = String("SubjectAccessReview")
	case *TokenReviewV1:
		v.ApiVersion = String("authentication.k8s.io/v1")
		v.Kind = String("TokenReview")
	case *ValidatingAdmissionPolicyBindingV1:
		v.ApiVersion = String("admissionregistration.k8s.io/v1")
		v.Kind = String("ValidatingAdmissionPolicyBinding")
	case *ValidatingAdmissionPolicyBindingV1alpha1:
		v.ApiVersion = String("admissionregistration.k8s.io/v1alpha1")
		v.Kind = String("ValidatingAdmissionPolicyBinding")
	case *ValidatingAdmissionPolicyBindingV1beta1:
		v.ApiVersion = String("admissionregistration.k8s.io/v1beta1")
		v.Kind = String("ValidatingAdmissionPolicyBinding")
	case *ValidatingAdmissionPolicyV1:
		v.ApiVersion = String("admissionregistration.k8s.io/v1")
		v.Kind = String("ValidatingAdmissionPolicy")
	case *ValidatingAdmissionPolicyV1alpha1:
		v.ApiVersion = String("admissionregistration.k8s.io/v1alpha1")
		v.Kind = String("ValidatingAdmissionPolicy")
	case *ValidatingAdmissionPolicyV1beta1:
		v.ApiVersion = String("admissionregistration.k8s.io/v1beta1")
		v.Kind = String("ValidatingAdmissionPolicy")
	case *ValidatingWebhookConfigurationV1:
		v.ApiVersion = String("admissionregistration.k8s.io/v1")
		v.Kind = String("ValidatingWebhookConfiguration")
	case *VolumeAttachmentV1:
		v.ApiVersion = String("storage.k8s.io/v1")
		v.Kind = String("VolumeAttachment")
	case *VolumeAttributesClassV1alpha1:
		v.ApiVersion = String("storage.k8s.io/v1alpha1")
		v.Kind = String("VolumeAttributesClass")
	}
}

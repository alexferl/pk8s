// Code generated; DO NOT EDIT.

package k8s

func SetAPIVersionAndKind(i any) {
	switch v := i.(type) {
	case *APIGroupV1:
		v.APIVersion = String("v1")
		v.Kind = String("APIGroup")
	case *APIServiceV1:
		v.APIVersion = String("apiregistration.k8s.io/v1")
		v.Kind = String("APIService")
	case *APIVersionsV1:
		v.APIVersion = String("v1")
		v.Kind = String("APIVersions")
	case *BindingV1:
		v.APIVersion = String("v1")
		v.Kind = String("Binding")
	case *CSIDriverV1:
		v.APIVersion = String("storage.k8s.io/v1")
		v.Kind = String("CSIDriver")
	case *CSINodeV1:
		v.APIVersion = String("storage.k8s.io/v1")
		v.Kind = String("CSINode")
	case *CSIStorageCapacityV1:
		v.APIVersion = String("storage.k8s.io/v1")
		v.Kind = String("CSIStorageCapacity")
	case *CertificateSigningRequestV1:
		v.APIVersion = String("certificates.k8s.io/v1")
		v.Kind = String("CertificateSigningRequest")
	case *ClusterRoleBindingV1:
		v.APIVersion = String("rbac.authorization.k8s.io/v1")
		v.Kind = String("ClusterRoleBinding")
	case *ClusterRoleV1:
		v.APIVersion = String("rbac.authorization.k8s.io/v1")
		v.Kind = String("ClusterRole")
	case *ClusterTrustBundleV1alpha1:
		v.APIVersion = String("certificates.k8s.io/v1alpha1")
		v.Kind = String("ClusterTrustBundle")
	case *ComponentStatusV1:
		v.APIVersion = String("v1")
		v.Kind = String("ComponentStatus")
	case *ConfigMapV1:
		v.APIVersion = String("v1")
		v.Kind = String("ConfigMap")
	case *ControllerRevisionV1:
		v.APIVersion = String("apps/v1")
		v.Kind = String("ControllerRevision")
	case *CronJobV1:
		v.APIVersion = String("batch/v1")
		v.Kind = String("CronJob")
	case *CustomResourceDefinitionV1:
		v.APIVersion = String("apiextensions.k8s.io/v1")
		v.Kind = String("CustomResourceDefinition")
	case *DaemonSetV1:
		v.APIVersion = String("apps/v1")
		v.Kind = String("DaemonSet")
	case *DeploymentV1:
		v.APIVersion = String("apps/v1")
		v.Kind = String("Deployment")
	case *DeviceClassV1alpha3:
		v.APIVersion = String("resource.k8s.io/v1alpha3")
		v.Kind = String("DeviceClass")
	case *EndpointSliceV1:
		v.APIVersion = String("discovery.k8s.io/v1")
		v.Kind = String("EndpointSlice")
	case *EndpointsV1:
		v.APIVersion = String("v1")
		v.Kind = String("Endpoints")
	case *EventV1:
		v.APIVersion = String("events.k8s.io/v1")
		v.Kind = String("Event")
	case *EvictionV1:
		v.APIVersion = String("policy/v1")
		v.Kind = String("Eviction")
	case *FlowSchemaV1:
		v.APIVersion = String("flowcontrol.apiserver.k8s.io/v1")
		v.Kind = String("FlowSchema")
	case *FlowSchemaV1beta3:
		v.APIVersion = String("flowcontrol.apiserver.k8s.io/v1beta3")
		v.Kind = String("FlowSchema")
	case *HorizontalPodAutoscalerV1:
		v.APIVersion = String("autoscaling/v1")
		v.Kind = String("HorizontalPodAutoscaler")
	case *HorizontalPodAutoscalerV2:
		v.APIVersion = String("autoscaling/v2")
		v.Kind = String("HorizontalPodAutoscaler")
	case *IPAddressV1beta1:
		v.APIVersion = String("networking.k8s.io/v1beta1")
		v.Kind = String("IPAddress")
	case *IngressClassV1:
		v.APIVersion = String("networking.k8s.io/v1")
		v.Kind = String("IngressClass")
	case *IngressV1:
		v.APIVersion = String("networking.k8s.io/v1")
		v.Kind = String("Ingress")
	case *JobV1:
		v.APIVersion = String("batch/v1")
		v.Kind = String("Job")
	case *LeaseCandidateV1alpha1:
		v.APIVersion = String("coordination.k8s.io/v1alpha1")
		v.Kind = String("LeaseCandidate")
	case *LeaseV1:
		v.APIVersion = String("coordination.k8s.io/v1")
		v.Kind = String("Lease")
	case *LimitRangeV1:
		v.APIVersion = String("v1")
		v.Kind = String("LimitRange")
	case *LocalSubjectAccessReviewV1:
		v.APIVersion = String("authorization.k8s.io/v1")
		v.Kind = String("LocalSubjectAccessReview")
	case *MutatingWebhookConfigurationV1:
		v.APIVersion = String("admissionregistration.k8s.io/v1")
		v.Kind = String("MutatingWebhookConfiguration")
	case *NamespaceV1:
		v.APIVersion = String("v1")
		v.Kind = String("Namespace")
	case *NetworkPolicyV1:
		v.APIVersion = String("networking.k8s.io/v1")
		v.Kind = String("NetworkPolicy")
	case *NodeV1:
		v.APIVersion = String("v1")
		v.Kind = String("Node")
	case *PersistentVolumeClaimV1:
		v.APIVersion = String("v1")
		v.Kind = String("PersistentVolumeClaim")
	case *PersistentVolumeV1:
		v.APIVersion = String("v1")
		v.Kind = String("PersistentVolume")
	case *PodDisruptionBudgetV1:
		v.APIVersion = String("policy/v1")
		v.Kind = String("PodDisruptionBudget")
	case *PodSchedulingContextV1alpha3:
		v.APIVersion = String("resource.k8s.io/v1alpha3")
		v.Kind = String("PodSchedulingContext")
	case *PodTemplateV1:
		v.APIVersion = String("v1")
		v.Kind = String("PodTemplate")
	case *PodV1:
		v.APIVersion = String("v1")
		v.Kind = String("Pod")
	case *PriorityClassV1:
		v.APIVersion = String("scheduling.k8s.io/v1")
		v.Kind = String("PriorityClass")
	case *PriorityLevelConfigurationV1:
		v.APIVersion = String("flowcontrol.apiserver.k8s.io/v1")
		v.Kind = String("PriorityLevelConfiguration")
	case *PriorityLevelConfigurationV1beta3:
		v.APIVersion = String("flowcontrol.apiserver.k8s.io/v1beta3")
		v.Kind = String("PriorityLevelConfiguration")
	case *ReplicaSetV1:
		v.APIVersion = String("apps/v1")
		v.Kind = String("ReplicaSet")
	case *ReplicationControllerV1:
		v.APIVersion = String("v1")
		v.Kind = String("ReplicationController")
	case *ResourceClaimTemplateV1alpha3:
		v.APIVersion = String("resource.k8s.io/v1alpha3")
		v.Kind = String("ResourceClaimTemplate")
	case *ResourceClaimV1alpha3:
		v.APIVersion = String("resource.k8s.io/v1alpha3")
		v.Kind = String("ResourceClaim")
	case *ResourceQuotaV1:
		v.APIVersion = String("v1")
		v.Kind = String("ResourceQuota")
	case *ResourceSliceV1alpha3:
		v.APIVersion = String("resource.k8s.io/v1alpha3")
		v.Kind = String("ResourceSlice")
	case *RoleBindingV1:
		v.APIVersion = String("rbac.authorization.k8s.io/v1")
		v.Kind = String("RoleBinding")
	case *RoleV1:
		v.APIVersion = String("rbac.authorization.k8s.io/v1")
		v.Kind = String("Role")
	case *RuntimeClassV1:
		v.APIVersion = String("node.k8s.io/v1")
		v.Kind = String("RuntimeClass")
	case *ScaleV1:
		v.APIVersion = String("autoscaling/v1")
		v.Kind = String("Scale")
	case *SecretV1:
		v.APIVersion = String("v1")
		v.Kind = String("Secret")
	case *SelfSubjectAccessReviewV1:
		v.APIVersion = String("authorization.k8s.io/v1")
		v.Kind = String("SelfSubjectAccessReview")
	case *SelfSubjectReviewV1:
		v.APIVersion = String("authentication.k8s.io/v1")
		v.Kind = String("SelfSubjectReview")
	case *SelfSubjectReviewV1alpha1:
		v.APIVersion = String("authentication.k8s.io/v1alpha1")
		v.Kind = String("SelfSubjectReview")
	case *SelfSubjectReviewV1beta1:
		v.APIVersion = String("authentication.k8s.io/v1beta1")
		v.Kind = String("SelfSubjectReview")
	case *SelfSubjectRulesReviewV1:
		v.APIVersion = String("authorization.k8s.io/v1")
		v.Kind = String("SelfSubjectRulesReview")
	case *ServiceAccountV1:
		v.APIVersion = String("v1")
		v.Kind = String("ServiceAccount")
	case *ServiceCIDRV1beta1:
		v.APIVersion = String("networking.k8s.io/v1beta1")
		v.Kind = String("ServiceCIDR")
	case *ServiceV1:
		v.APIVersion = String("v1")
		v.Kind = String("Service")
	case *StatefulSetV1:
		v.APIVersion = String("apps/v1")
		v.Kind = String("StatefulSet")
	case *StatusV1:
		v.APIVersion = String("v1")
		v.Kind = String("Status")
	case *StorageClassV1:
		v.APIVersion = String("storage.k8s.io/v1")
		v.Kind = String("StorageClass")
	case *StorageVersionMigrationV1alpha1:
		v.APIVersion = String("storagemigration.k8s.io/v1alpha1")
		v.Kind = String("StorageVersionMigration")
	case *StorageVersionV1alpha1:
		v.APIVersion = String("internal.apiserver.k8s.io/v1alpha1")
		v.Kind = String("StorageVersion")
	case *SubjectAccessReviewV1:
		v.APIVersion = String("authorization.k8s.io/v1")
		v.Kind = String("SubjectAccessReview")
	case *TokenReviewV1:
		v.APIVersion = String("authentication.k8s.io/v1")
		v.Kind = String("TokenReview")
	case *ValidatingAdmissionPolicyBindingV1:
		v.APIVersion = String("admissionregistration.k8s.io/v1")
		v.Kind = String("ValidatingAdmissionPolicyBinding")
	case *ValidatingAdmissionPolicyBindingV1alpha1:
		v.APIVersion = String("admissionregistration.k8s.io/v1alpha1")
		v.Kind = String("ValidatingAdmissionPolicyBinding")
	case *ValidatingAdmissionPolicyBindingV1beta1:
		v.APIVersion = String("admissionregistration.k8s.io/v1beta1")
		v.Kind = String("ValidatingAdmissionPolicyBinding")
	case *ValidatingAdmissionPolicyV1:
		v.APIVersion = String("admissionregistration.k8s.io/v1")
		v.Kind = String("ValidatingAdmissionPolicy")
	case *ValidatingAdmissionPolicyV1alpha1:
		v.APIVersion = String("admissionregistration.k8s.io/v1alpha1")
		v.Kind = String("ValidatingAdmissionPolicy")
	case *ValidatingAdmissionPolicyV1beta1:
		v.APIVersion = String("admissionregistration.k8s.io/v1beta1")
		v.Kind = String("ValidatingAdmissionPolicy")
	case *ValidatingWebhookConfigurationV1:
		v.APIVersion = String("admissionregistration.k8s.io/v1")
		v.Kind = String("ValidatingWebhookConfiguration")
	case *VolumeAttachmentV1:
		v.APIVersion = String("storage.k8s.io/v1")
		v.Kind = String("VolumeAttachment")
	case *VolumeAttributesClassV1alpha1:
		v.APIVersion = String("storage.k8s.io/v1alpha1")
		v.Kind = String("VolumeAttributesClass")
	case *VolumeAttributesClassV1beta1:
		v.APIVersion = String("storage.k8s.io/v1beta1")
		v.Kind = String("VolumeAttributesClass")
	}
}

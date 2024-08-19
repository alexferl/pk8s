package k8s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetApiVersionAndKind(t *testing.T) {
	tests := []struct {
		name       string
		input      any
		wantApiVer string
		wantKind   string
	}{
		{
			name:       "APIGroupV1",
			input:      &APIGroupV1{},
			wantApiVer: "v1",
			wantKind:   "APIGroup",
		},
		{
			name:       "APIServiceV1",
			input:      &APIServiceV1{},
			wantApiVer: "apiregistration.k8s.io/v1",
			wantKind:   "APIService",
		},
		{
			name:       "APIVersionsV1",
			input:      &APIVersionsV1{},
			wantApiVer: "v1",
			wantKind:   "APIVersions",
		},
		{
			name:       "BindingV1",
			input:      &BindingV1{},
			wantApiVer: "v1",
			wantKind:   "Binding",
		},
		{
			name:       "CSIDriverV1",
			input:      &CSIDriverV1{},
			wantApiVer: "storage.k8s.io/v1",
			wantKind:   "CSIDriver",
		},
		{
			name:       "CSINodeV1",
			input:      &CSINodeV1{},
			wantApiVer: "storage.k8s.io/v1",
			wantKind:   "CSINode",
		},
		{
			name:       "CSIStorageCapacityV1",
			input:      &CSIStorageCapacityV1{},
			wantApiVer: "storage.k8s.io/v1",
			wantKind:   "CSIStorageCapacity",
		},
		{
			name:       "CertificateSigningRequestV1",
			input:      &CertificateSigningRequestV1{},
			wantApiVer: "certificates.k8s.io/v1",
			wantKind:   "CertificateSigningRequest",
		},
		{
			name:       "ClusterRoleBindingV1",
			input:      &ClusterRoleBindingV1{},
			wantApiVer: "rbac.authorization.k8s.io/v1",
			wantKind:   "ClusterRoleBinding",
		},
		{
			name:       "ClusterRoleV1",
			input:      &ClusterRoleV1{},
			wantApiVer: "rbac.authorization.k8s.io/v1",
			wantKind:   "ClusterRole",
		},
		{
			name:       "ClusterTrustBundleV1alpha1",
			input:      &ClusterTrustBundleV1alpha1{},
			wantApiVer: "certificates.k8s.io/v1alpha1",
			wantKind:   "ClusterTrustBundle",
		},
		{
			name:       "ComponentStatusV1",
			input:      &ComponentStatusV1{},
			wantApiVer: "v1",
			wantKind:   "ComponentStatus",
		},
		{
			name:       "ConfigMapV1",
			input:      &ConfigMapV1{},
			wantApiVer: "v1",
			wantKind:   "ConfigMap",
		},
		{
			name:       "ControllerRevisionV1",
			input:      &ControllerRevisionV1{},
			wantApiVer: "apps/v1",
			wantKind:   "ControllerRevision",
		},
		{
			name:       "CronJobV1",
			input:      &CronJobV1{},
			wantApiVer: "batch/v1",
			wantKind:   "CronJob",
		},
		{
			name:       "CustomResourceDefinitionV1",
			input:      &CustomResourceDefinitionV1{},
			wantApiVer: "apiextensions.k8s.io/v1",
			wantKind:   "CustomResourceDefinition",
		},
		{
			name:       "DaemonSetV1",
			input:      &DaemonSetV1{},
			wantApiVer: "apps/v1",
			wantKind:   "DaemonSet",
		},
		{
			name:       "DeploymentV1",
			input:      &DeploymentV1{},
			wantApiVer: "apps/v1",
			wantKind:   "Deployment",
		},
		{
			name:       "EndpointSliceV1",
			input:      &EndpointSliceV1{},
			wantApiVer: "discovery.k8s.io/v1",
			wantKind:   "EndpointSlice",
		},
		{
			name:       "EndpointsV1",
			input:      &EndpointsV1{},
			wantApiVer: "v1",
			wantKind:   "Endpoints",
		},
		{
			name:       "EventV1",
			input:      &EventV1{},
			wantApiVer: "events.k8s.io/v1",
			wantKind:   "Event",
		},
		{
			name:       "EvictionV1",
			input:      &EvictionV1{},
			wantApiVer: "policy/v1",
			wantKind:   "Eviction",
		},
		{
			name:       "FlowSchemaV1",
			input:      &FlowSchemaV1{},
			wantApiVer: "flowcontrol.apiserver.k8s.io/v1",
			wantKind:   "FlowSchema",
		},
		{
			name:       "FlowSchemaV1beta3",
			input:      &FlowSchemaV1beta3{},
			wantApiVer: "flowcontrol.apiserver.k8s.io/v1beta3",
			wantKind:   "FlowSchema",
		},
		{
			name:       "HorizontalPodAutoscalerV1",
			input:      &HorizontalPodAutoscalerV1{},
			wantApiVer: "autoscaling/v1",
			wantKind:   "HorizontalPodAutoscaler",
		},
		{
			name:       "HorizontalPodAutoscalerV2",
			input:      &HorizontalPodAutoscalerV2{},
			wantApiVer: "autoscaling/v2",
			wantKind:   "HorizontalPodAutoscaler",
		},
		{
			name:       "IPAddressV1alpha1",
			input:      &IPAddressV1alpha1{},
			wantApiVer: "networking.k8s.io/v1alpha1",
			wantKind:   "IPAddress",
		},
		{
			name:       "IngressClassV1",
			input:      &IngressClassV1{},
			wantApiVer: "networking.k8s.io/v1",
			wantKind:   "IngressClass",
		},
		{
			name:       "IngressV1",
			input:      &IngressV1{},
			wantApiVer: "networking.k8s.io/v1",
			wantKind:   "Ingress",
		},
		{
			name:       "JobV1",
			input:      &JobV1{},
			wantApiVer: "batch/v1",
			wantKind:   "Job",
		},
		{
			name:       "LeaseV1",
			input:      &LeaseV1{},
			wantApiVer: "coordination.k8s.io/v1",
			wantKind:   "Lease",
		},
		{
			name:       "LimitRangeV1",
			input:      &LimitRangeV1{},
			wantApiVer: "v1",
			wantKind:   "LimitRange",
		},
		{
			name:       "LocalSubjectAccessReviewV1",
			input:      &LocalSubjectAccessReviewV1{},
			wantApiVer: "authorization.k8s.io/v1",
			wantKind:   "LocalSubjectAccessReview",
		},
		{
			name:       "MutatingWebhookConfigurationV1",
			input:      &MutatingWebhookConfigurationV1{},
			wantApiVer: "admissionregistration.k8s.io/v1",
			wantKind:   "MutatingWebhookConfiguration",
		},
		{
			name:       "NamespaceV1",
			input:      &NamespaceV1{},
			wantApiVer: "v1",
			wantKind:   "Namespace",
		},
		{
			name:       "NetworkPolicyV1",
			input:      &NetworkPolicyV1{},
			wantApiVer: "networking.k8s.io/v1",
			wantKind:   "NetworkPolicy",
		},
		{
			name:       "NodeV1",
			input:      &NodeV1{},
			wantApiVer: "v1",
			wantKind:   "Node",
		},
		{
			name:       "PersistentVolumeClaimV1",
			input:      &PersistentVolumeClaimV1{},
			wantApiVer: "v1",
			wantKind:   "PersistentVolumeClaim",
		},
		{
			name:       "PersistentVolumeV1",
			input:      &PersistentVolumeV1{},
			wantApiVer: "v1",
			wantKind:   "PersistentVolume",
		},
		{
			name:       "PodDisruptionBudgetV1",
			input:      &PodDisruptionBudgetV1{},
			wantApiVer: "policy/v1",
			wantKind:   "PodDisruptionBudget",
		},
		{
			name:       "PodSchedulingContextV1alpha2",
			input:      &PodSchedulingContextV1alpha2{},
			wantApiVer: "resource.k8s.io/v1alpha2",
			wantKind:   "PodSchedulingContext",
		},
		{
			name:       "PodTemplateV1",
			input:      &PodTemplateV1{},
			wantApiVer: "v1",
			wantKind:   "PodTemplate",
		},
		{
			name:       "PodV1",
			input:      &PodV1{},
			wantApiVer: "v1",
			wantKind:   "Pod",
		},
		{
			name:       "PriorityClassV1",
			input:      &PriorityClassV1{},
			wantApiVer: "scheduling.k8s.io/v1",
			wantKind:   "PriorityClass",
		},
		{
			name:       "PriorityLevelConfigurationV1",
			input:      &PriorityLevelConfigurationV1{},
			wantApiVer: "flowcontrol.apiserver.k8s.io/v1",
			wantKind:   "PriorityLevelConfiguration",
		},
		{
			name:       "PriorityLevelConfigurationV1beta3",
			input:      &PriorityLevelConfigurationV1beta3{},
			wantApiVer: "flowcontrol.apiserver.k8s.io/v1beta3",
			wantKind:   "PriorityLevelConfiguration",
		},
		{
			name:       "ReplicaSetV1",
			input:      &ReplicaSetV1{},
			wantApiVer: "apps/v1",
			wantKind:   "ReplicaSet",
		},
		{
			name:       "ReplicationControllerV1",
			input:      &ReplicationControllerV1{},
			wantApiVer: "v1",
			wantKind:   "ReplicationController",
		},
		{
			name:       "ResourceClaimParametersV1alpha2",
			input:      &ResourceClaimParametersV1alpha2{},
			wantApiVer: "resource.k8s.io/v1alpha2",
			wantKind:   "ResourceClaimParameters",
		},
		{
			name:       "ResourceClaimTemplateV1alpha2",
			input:      &ResourceClaimTemplateV1alpha2{},
			wantApiVer: "resource.k8s.io/v1alpha2",
			wantKind:   "ResourceClaimTemplate",
		},
		{
			name:       "ResourceClaimV1alpha2",
			input:      &ResourceClaimV1alpha2{},
			wantApiVer: "resource.k8s.io/v1alpha2",
			wantKind:   "ResourceClaim",
		},
		{
			name:       "ResourceClassParametersV1alpha2",
			input:      &ResourceClassParametersV1alpha2{},
			wantApiVer: "resource.k8s.io/v1alpha2",
			wantKind:   "ResourceClassParameters",
		},
		{
			name:       "ResourceClassV1alpha2",
			input:      &ResourceClassV1alpha2{},
			wantApiVer: "resource.k8s.io/v1alpha2",
			wantKind:   "ResourceClass",
		},
		{
			name:       "ResourceQuotaV1",
			input:      &ResourceQuotaV1{},
			wantApiVer: "v1",
			wantKind:   "ResourceQuota",
		},
		{
			name:       "ResourceSliceV1alpha2",
			input:      &ResourceSliceV1alpha2{},
			wantApiVer: "resource.k8s.io/v1alpha2",
			wantKind:   "ResourceSlice",
		},
		{
			name:       "RoleBindingV1",
			input:      &RoleBindingV1{},
			wantApiVer: "rbac.authorization.k8s.io/v1",
			wantKind:   "RoleBinding",
		},
		{
			name:       "RoleV1",
			input:      &RoleV1{},
			wantApiVer: "rbac.authorization.k8s.io/v1",
			wantKind:   "Role",
		},
		{
			name:       "RuntimeClassV1",
			input:      &RuntimeClassV1{},
			wantApiVer: "node.k8s.io/v1",
			wantKind:   "RuntimeClass",
		},
		{
			name:       "ScaleV1",
			input:      &ScaleV1{},
			wantApiVer: "autoscaling/v1",
			wantKind:   "Scale",
		},
		{
			name:       "SecretV1",
			input:      &SecretV1{},
			wantApiVer: "v1",
			wantKind:   "Secret",
		},
		{
			name:       "SelfSubjectAccessReviewV1",
			input:      &SelfSubjectAccessReviewV1{},
			wantApiVer: "authorization.k8s.io/v1",
			wantKind:   "SelfSubjectAccessReview",
		},
		{
			name:       "SelfSubjectReviewV1",
			input:      &SelfSubjectReviewV1{},
			wantApiVer: "authentication.k8s.io/v1",
			wantKind:   "SelfSubjectReview",
		},
		{
			name:       "SelfSubjectReviewV1alpha1",
			input:      &SelfSubjectReviewV1alpha1{},
			wantApiVer: "authentication.k8s.io/v1alpha1",
			wantKind:   "SelfSubjectReview",
		},
		{
			name:       "SelfSubjectReviewV1beta1",
			input:      &SelfSubjectReviewV1beta1{},
			wantApiVer: "authentication.k8s.io/v1beta1",
			wantKind:   "SelfSubjectReview",
		},
		{
			name:       "SelfSubjectRulesReviewV1",
			input:      &SelfSubjectRulesReviewV1{},
			wantApiVer: "authorization.k8s.io/v1",
			wantKind:   "SelfSubjectRulesReview",
		},
		{
			name:       "ServiceAccountV1",
			input:      &ServiceAccountV1{},
			wantApiVer: "v1",
			wantKind:   "ServiceAccount",
		},
		{
			name:       "ServiceCIDRV1alpha1",
			input:      &ServiceCIDRV1alpha1{},
			wantApiVer: "networking.k8s.io/v1alpha1",
			wantKind:   "ServiceCIDR",
		},
		{
			name:       "ServiceV1",
			input:      &ServiceV1{},
			wantApiVer: "v1",
			wantKind:   "Service",
		},
		{
			name:       "StatefulSetV1",
			input:      &StatefulSetV1{},
			wantApiVer: "apps/v1",
			wantKind:   "StatefulSet",
		},
		{
			name:       "StatusV1",
			input:      &StatusV1{},
			wantApiVer: "v1",
			wantKind:   "Status",
		},
		{
			name:       "StorageClassV1",
			input:      &StorageClassV1{},
			wantApiVer: "storage.k8s.io/v1",
			wantKind:   "StorageClass",
		},
		{
			name:       "StorageVersionMigrationV1alpha1",
			input:      &StorageVersionMigrationV1alpha1{},
			wantApiVer: "storagemigration.k8s.io/v1alpha1",
			wantKind:   "StorageVersionMigration",
		},
		{
			name:       "StorageVersionV1alpha1",
			input:      &StorageVersionV1alpha1{},
			wantApiVer: "internal.apiserver.k8s.io/v1alpha1",
			wantKind:   "StorageVersion",
		},
		{
			name:       "SubjectAccessReviewV1",
			input:      &SubjectAccessReviewV1{},
			wantApiVer: "authorization.k8s.io/v1",
			wantKind:   "SubjectAccessReview",
		},
		{
			name:       "TokenReviewV1",
			input:      &TokenReviewV1{},
			wantApiVer: "authentication.k8s.io/v1",
			wantKind:   "TokenReview",
		},
		{
			name:       "ValidatingAdmissionPolicyBindingV1",
			input:      &ValidatingAdmissionPolicyBindingV1{},
			wantApiVer: "admissionregistration.k8s.io/v1",
			wantKind:   "ValidatingAdmissionPolicyBinding",
		},
		{
			name:       "ValidatingAdmissionPolicyBindingV1alpha1",
			input:      &ValidatingAdmissionPolicyBindingV1alpha1{},
			wantApiVer: "admissionregistration.k8s.io/v1alpha1",
			wantKind:   "ValidatingAdmissionPolicyBinding",
		},
		{
			name:       "ValidatingAdmissionPolicyBindingV1beta1",
			input:      &ValidatingAdmissionPolicyBindingV1beta1{},
			wantApiVer: "admissionregistration.k8s.io/v1beta1",
			wantKind:   "ValidatingAdmissionPolicyBinding",
		},
		{
			name:       "ValidatingAdmissionPolicyV1",
			input:      &ValidatingAdmissionPolicyV1{},
			wantApiVer: "admissionregistration.k8s.io/v1",
			wantKind:   "ValidatingAdmissionPolicy",
		},
		{
			name:       "ValidatingAdmissionPolicyV1alpha1",
			input:      &ValidatingAdmissionPolicyV1alpha1{},
			wantApiVer: "admissionregistration.k8s.io/v1alpha1",
			wantKind:   "ValidatingAdmissionPolicy",
		},
		{
			name:       "ValidatingAdmissionPolicyV1beta1",
			input:      &ValidatingAdmissionPolicyV1beta1{},
			wantApiVer: "admissionregistration.k8s.io/v1beta1",
			wantKind:   "ValidatingAdmissionPolicy",
		},
		{
			name:       "ValidatingWebhookConfigurationV1",
			input:      &ValidatingWebhookConfigurationV1{},
			wantApiVer: "admissionregistration.k8s.io/v1",
			wantKind:   "ValidatingWebhookConfiguration",
		},
		{
			name:       "VolumeAttachmentV1",
			input:      &VolumeAttachmentV1{},
			wantApiVer: "storage.k8s.io/v1",
			wantKind:   "VolumeAttachment",
		},
		{
			name:       "VolumeAttributesClassV1alpha1",
			input:      &VolumeAttributesClassV1alpha1{},
			wantApiVer: "storage.k8s.io/v1alpha1",
			wantKind:   "VolumeAttributesClass",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetApiVersionAndKind(tt.input)

			switch obj := tt.input.(type) {
			case *APIGroupV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *APIServiceV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *APIVersionsV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *BindingV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *CSIDriverV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *CSINodeV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *CSIStorageCapacityV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *CertificateSigningRequestV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ClusterRoleBindingV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ClusterRoleV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ClusterTrustBundleV1alpha1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ComponentStatusV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ConfigMapV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ControllerRevisionV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *CronJobV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *CustomResourceDefinitionV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *DaemonSetV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *DeploymentV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *EndpointSliceV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *EndpointsV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *EventV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *EvictionV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *FlowSchemaV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *FlowSchemaV1beta3:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *HorizontalPodAutoscalerV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *HorizontalPodAutoscalerV2:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *IPAddressV1alpha1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *IngressClassV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *IngressV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *JobV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *LeaseV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *LimitRangeV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *LocalSubjectAccessReviewV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *MutatingWebhookConfigurationV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *NamespaceV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *NetworkPolicyV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *NodeV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *PersistentVolumeClaimV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *PersistentVolumeV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *PodDisruptionBudgetV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *PodSchedulingContextV1alpha2:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *PodTemplateV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *PodV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *PriorityClassV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *PriorityLevelConfigurationV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *PriorityLevelConfigurationV1beta3:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ReplicaSetV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ReplicationControllerV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ResourceClaimParametersV1alpha2:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ResourceClaimTemplateV1alpha2:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ResourceClaimV1alpha2:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ResourceClassParametersV1alpha2:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ResourceClassV1alpha2:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ResourceQuotaV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ResourceSliceV1alpha2:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *RoleBindingV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *RoleV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *RuntimeClassV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ScaleV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *SecretV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *SelfSubjectAccessReviewV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *SelfSubjectReviewV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *SelfSubjectReviewV1alpha1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *SelfSubjectReviewV1beta1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *SelfSubjectRulesReviewV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ServiceAccountV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ServiceCIDRV1alpha1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ServiceV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *StatefulSetV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *StatusV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *StorageClassV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *StorageVersionMigrationV1alpha1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *StorageVersionV1alpha1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *SubjectAccessReviewV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *TokenReviewV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ValidatingAdmissionPolicyBindingV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ValidatingAdmissionPolicyBindingV1alpha1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ValidatingAdmissionPolicyBindingV1beta1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ValidatingAdmissionPolicyV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ValidatingAdmissionPolicyV1alpha1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ValidatingAdmissionPolicyV1beta1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *ValidatingWebhookConfigurationV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *VolumeAttachmentV1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			case *VolumeAttributesClassV1alpha1:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
			}
		})
	}
}

package handler

import (
	v1alpha1 "github.com/zdq0394/consul-operator/pkg/apis/consul/v1alpha1"
	"github.com/zdq0394/consul-operator/service/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConsulManager manage resource instance in kubernetes cluster
type ConsulManager interface {
	EnsureConsulConfigMap(rc *v1alpha1.Consul, labels map[string]string, ownerRefs []metav1.OwnerReference) error
	EnsureConsulStatefulset(rc *v1alpha1.Consul, lables map[string]string, ownerRefs []metav1.OwnerReference) error
	EnsureConsulHeadlessService(rc *v1alpha1.Consul, labels map[string]string, ownerRefs []metav1.OwnerReference) error
}

type consulKubeClusterManager struct {
	ClusterDomain string
	K8SService    k8s.Services
}

// NewConsulManager new consul cluster manager.
func NewConsulManager(s k8s.Services, cd string) ConsulManager {
	return &consulKubeClusterManager{
		K8SService:    s,
		ClusterDomain: cd,
	}
}

func (s *consulKubeClusterManager) EnsureConsulConfigMap(rc *v1alpha1.Consul, labels map[string]string, ownerRefs []metav1.OwnerReference) error {
	configMap := generateConfigMap(rc, labels, ownerRefs, s.ClusterDomain)
	return s.K8SService.CreateOrUpdateConfigMap(rc.Namespace, configMap)
}

func (s *consulKubeClusterManager) EnsureConsulStatefulset(rc *v1alpha1.Consul, labels map[string]string, ownerRefs []metav1.OwnerReference) error {
	ss := generateConsulStatefulset(rc, labels, ownerRefs)
	return s.K8SService.CreateOrUpdateStatefulSet(rc.Namespace, ss)
}

func (s *consulKubeClusterManager) EnsureConsulHeadlessService(rc *v1alpha1.Consul, labels map[string]string, ownerRefs []metav1.OwnerReference) error {
	svc := generateConsulHeadlessService(rc, labels, ownerRefs)
	return s.K8SService.CreateIfNotExistsService(rc.Namespace, svc)
}

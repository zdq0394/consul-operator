package consul

import (
	"github.com/zdq0394/consul-operator/pkg/apis/consul/v1alpha1"
	"github.com/zdq0394/consul-operator/service/k8s"
	"github.com/zdq0394/k8soperator/pkg/operator/crd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type ConsulCRD struct {
	service k8s.Services
}

// NewConsulCRD create an instance of ConsulCRD
func NewConsulCRD(service k8s.Services) *ConsulCRD {
	return &ConsulCRD{
		service: service,
	}
}

// Initialize ensure ConsulCRD created in k8s.
func (s *ConsulCRD) Initialize() error {
	crdConf := crd.Conf{
		Kind:       v1alpha1.RCKind,
		NamePlural: v1alpha1.RCNamePlural,
		Group:      v1alpha1.SchemeGroupVersion.Group,
		Version:    v1alpha1.SchemeGroupVersion.Version,
		Scope:      v1alpha1.RCScope,
		Categories: []string{"all"},
	}
	return s.service.EnsureCRD(crdConf)
}

// GetListerWatcher get the listwatcher of ConsulCRD.
func (s *ConsulCRD) GetListerWatcher() cache.ListerWatcher {
	return &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return s.service.ListConsuls("", options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return s.service.WatchConsuls("", options)
		},
	}
}

// GetObject get the Consul Object
func (s *ConsulCRD) GetObject() runtime.Object {
	return &v1alpha1.Consul{}
}

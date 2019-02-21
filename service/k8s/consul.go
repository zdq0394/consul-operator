package k8s

import (
	consulv1alpha1 "github.com/zdq0394/consul-operator/pkg/apis/consul/v1alpha1"
	clientset "github.com/zdq0394/consul-operator/pkg/client/clientset/versioned"
	"github.com/zdq0394/k8soperator/pkg/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

// Consul the Consul service that knows how to interact with k8s to get them
type Consul interface {
	ListConsuls(namespace string, opts metav1.ListOptions) (*consulv1alpha1.ConsulList, error)
	WatchConsuls(namespace string, opts metav1.ListOptions) (watch.Interface, error)
}

// ConsulService is the RedisCluster service implementation using API calls to kubernetes.
type ConsulService struct {
	crdClient clientset.Interface
	logger    log.Logger
}

// NewConsulService returns a new Workspace KubeService.
func NewConsulService(crdcli clientset.Interface, logger log.Logger) *ConsulService {
	logger = logger.With("service", "k8s.consuls")
	return &ConsulService{
		crdClient: crdcli,
		logger:    logger,
	}
}

// ListConsuls ...
func (r *ConsulService) ListConsuls(namespace string, opts metav1.ListOptions) (*consulv1alpha1.ConsulList, error) {
	return r.crdClient.Consul().Consuls(namespace).List(opts)
}

// WatchConsuls ...
func (r *ConsulService) WatchConsuls(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return r.crdClient.Consul().Consuls(namespace).Watch(opts)
}

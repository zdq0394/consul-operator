package k8s

import (
	apiextensionscli "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"

	versionedClientset "github.com/zdq0394/consul-operator/pkg/client/clientset/versioned"
	"github.com/zdq0394/consul-operator/pkg/log"
)

// Services is the K8s service entrypoint.
type Services interface {
	CRD
	ConfigMap
	Pod
	PodDisruptionBudget
	Consul
	Service
	RBAC
	Deployment
	StatefulSet
}

type services struct {
	CRD
	ConfigMap
	Pod
	PodDisruptionBudget
	Consul
	Service
	RBAC
	Deployment
	StatefulSet
}

// New returns a new Kubernetes service.
func New(kubecli kubernetes.Interface, crdcli versionedClientset.Interface, apiextcli apiextensionscli.Interface, logger log.Logger) Services {
	return &services{
		CRD:                 NewCRDService(apiextcli, logger),
		ConfigMap:           NewConfigMapService(kubecli, logger),
		Pod:                 NewPodService(kubecli, logger),
		PodDisruptionBudget: NewPodDisruptionBudgetService(kubecli, logger),
		Consul:              NewConsulService(crdcli, logger),
		Service:             NewServiceService(kubecli, logger),
		RBAC:                NewRBACService(kubecli, logger),
		Deployment:          NewDeploymentService(kubecli, logger),
		StatefulSet:         NewStatefulSetService(kubecli, logger),
	}
}

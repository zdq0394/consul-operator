package consul

import (
	mgr "github.com/zdq0394/consul-operator/operator/consul/handler"
	"github.com/zdq0394/consul-operator/pkg/k8sutil"
	"github.com/zdq0394/consul-operator/pkg/log"
	"github.com/zdq0394/consul-operator/service/k8s"
	"github.com/zdq0394/k8soperator/pkg/operator"
	"github.com/zdq0394/k8soperator/pkg/operator/controller"
)

type Config struct {
	Development       bool
	Kubeconfig        string
	ClusterDomain     string
	ConcurrentWorkers int
}

// Start the Operator
func Start(conf *Config) error {
	kubeClient, consulClient, aeClient, _ := k8sutil.CreateKubernetesClients(conf.Development, conf.Kubeconfig)

	logger := log.Base()
	k8sService := k8s.New(kubeClient, consulClient, aeClient, logger)

	crd := NewConsulCRD(k8sService)
	manager := mgr.NewConsulManager(k8sService, conf.ClusterDomain)
	handler := NewConsulHandler(nil, manager, logger)

	controllerCfg := &controller.Config{
		Name:              "Consul Controller",
		ConcurrentWorkers: conf.ConcurrentWorkers,
	}
	ctrl := controller.NewSimpleController(controllerCfg, crd, handler)
	optor := operator.NewSimpleOperator(crd, ctrl)
	stopC := make(chan struct{}, 0)
	optor.Run(stopC)
	return nil
}

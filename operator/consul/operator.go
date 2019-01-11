package consul

import (
	k8sclient "github.com/zdq0394/consul-operator/pkg/k8s"
	"github.com/zdq0394/consul-operator/pkg/log"
	k8service "github.com/zdq0394/consul-operator/service/k8s"
	"github.com/zdq0394/k8soperator/pkg/operator"
	"github.com/zdq0394/k8soperator/pkg/operator/controller"
)

type Config struct {
	Development       bool
	Kubeconfig        string
	ConcurrentWorkers int
}

// Start the Operator
func Start(conf *Config) error {
	kubeClient, consulClient, aeClient, _ := k8sclient.CreateKubernetesClients(conf.Development, conf.Kubeconfig)
	logger := log.Base()
	kubeService := k8service.New(kubeClient, consulClient, aeClient, logger)

	crd := NewConsulCRD(kubeService)

	handler := NewConsulHandler(nil)

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

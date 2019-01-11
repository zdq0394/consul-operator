package consul

import (
	"context"
	"fmt"

	"github.com/zdq0394/consul-operator/operator"
	mgr "github.com/zdq0394/consul-operator/operator/consul/handler"
	"github.com/zdq0394/consul-operator/pkg/apis/consul/v1alpha1"
	"github.com/zdq0394/consul-operator/pkg/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var defaultLabels = map[string]string{
	"Creator": "ConsulOperator",
}

const (
	ConsulLabelKey = "consul"
)

type ConsulHandler struct {
	Manager mgr.ConsulManager
	Labels  map[string]string
}

func NewConsulHandler(labels map[string]string, manager mgr.ConsulManager) *ConsulHandler {
	curLabels := operator.MergeLabels(defaultLabels, labels)
	return &ConsulHandler{
		Labels:  curLabels,
		Manager: manager,
	}
}

func (h *ConsulHandler) Add(ctx context.Context, obj runtime.Object) error {
	log.Infoln("Create Consul Here...")
	rc, ok := obj.(*v1alpha1.Consul)
	if !ok {
		return fmt.Errorf("Cannot handle Consul")
	}
	log.Infof("Handler Create Consul:%s/%s\n", rc.Namespace, rc.Name)
	oRefs := h.createOwnerReferences(rc)
	instanceLabels := h.generateInstanceLabels(rc)
	labels := operator.MergeLabels(h.Labels, rc.Labels, instanceLabels)
	return h.ensurePresent(rc, labels, oRefs)
}

func (h *ConsulHandler) Delete(ctx context.Context, key string) error {
	log.Infoln("Delete Consul Here:", key)
	// No need to do anything, it will be handled by the owner reference done
	// on the creation.
	return nil
}

func (h *ConsulHandler) createOwnerReferences(rc *v1alpha1.Consul) []metav1.OwnerReference {
	rcvk := v1alpha1.VersionKind(v1alpha1.RCKind)
	return []metav1.OwnerReference{
		*metav1.NewControllerRef(rc, rcvk),
	}
}

func (h *ConsulHandler) generateInstanceLabels(rc *v1alpha1.Consul) map[string]string {
	return map[string]string{
		ConsulLabelKey: rc.Name,
	}
}

func (h *ConsulHandler) ensurePresent(rc *v1alpha1.Consul,
	labels map[string]string, ownerRefs []metav1.OwnerReference) error {
	if err := h.Manager.EnsureConsulConfigMap(rc, labels, ownerRefs); err != nil {
		return err
	}
	if err := h.Manager.EnsureConsulHeadlessService(rc, labels, ownerRefs); err != nil {
		return err
	}
	if err := h.Manager.EnsureConsulStatefulset(rc, labels, ownerRefs); err != nil {
		return err
	}
	return nil
}

package consul

import (
	"context"
	"fmt"

	mgr "github.com/zdq0394/consul-operator/operator/consul/handler"
	"github.com/zdq0394/consul-operator/pkg/apis/consul/v1alpha1"
	"github.com/zdq0394/k8soperator/pkg/log"
	"github.com/zdq0394/k8soperator/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var defaultLabels = map[string]string{
	"Creator": "ConsulOperator",
}

const (
	// InstanceLabelKey its value is the name of consul.consul.zdq0394.io resouce.
	InstanceLabelKey = "ConsulInstance"
)

// Handler handles events of add/update/delete of Consul Resouce.
type Handler struct {
	Manager mgr.ConsulManager
	Labels  map[string]string
	logger  log.Logger
}

// NewConsulHandler return a Handler instance.
func NewConsulHandler(labels map[string]string, manager mgr.ConsulManager, logger log.Logger) *Handler {
	curLabels := util.MergeLabels(defaultLabels, labels)
	return &Handler{
		Labels:  curLabels,
		Manager: manager,
		logger:  logger,
	}
}

// Add will be called when a Consul resource added/updated.
func (h *Handler) Add(ctx context.Context, obj runtime.Object) error {
	h.logger.Infoln("Create Consul Here...")
	rc, ok := obj.(*v1alpha1.Consul)
	if !ok {
		return fmt.Errorf("Cannot handle Consul")
	}
	h.logger.Infof("Handler Create Consul:%s/%s\n", rc.Namespace, rc.Name)
	oRefs := h.createOwnerReferences(rc)
	instanceLabels := h.generateInstanceLabels(rc)
	labels := util.MergeLabels(h.Labels, rc.Labels, instanceLabels)
	return h.ensurePresent(rc, labels, oRefs)
}

// Delete will be called when a Consul resource deleted.
func (h *Handler) Delete(ctx context.Context, key string) error {
	h.logger.Infoln("Delete Consul Here:", key)
	// No need to do anything, Kubernetes will handle it by the owner reference done on the creation.
	return nil
}

func (h *Handler) createOwnerReferences(rc *v1alpha1.Consul) []metav1.OwnerReference {
	rcvk := v1alpha1.VersionKind(v1alpha1.RCKind)
	return []metav1.OwnerReference{
		*metav1.NewControllerRef(rc, rcvk),
	}
}

func (h *Handler) generateInstanceLabels(rc *v1alpha1.Consul) map[string]string {
	return map[string]string{
		InstanceLabelKey: rc.Name,
	}
}

func (h *Handler) ensurePresent(rc *v1alpha1.Consul,
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

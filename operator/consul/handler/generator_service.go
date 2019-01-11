package handler

import (
	"github.com/zdq0394/consul-operator/pkg/apis/consul/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func generateConsulHeadlessService(rc *v1alpha1.Consul,
	labels map[string]string, ownerRefs []metav1.OwnerReference) *corev1.Service {

	name := generateName(svcHeadlessNamePrefix, rc.Name)
	namespace := rc.Namespace

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:            name,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: ownerRefs,
		},
		Spec: corev1.ServiceSpec{
			Type:      corev1.ServiceTypeClusterIP,
			ClusterIP: corev1.ClusterIPNone,
			Ports: []corev1.ServicePort{
				{
					Port:     consulHeadlessPort,
					Protocol: corev1.ProtocolTCP,
					Name:     consulHeadlessPortName,
				},
			},
			Selector: labels,
		},
	}
}

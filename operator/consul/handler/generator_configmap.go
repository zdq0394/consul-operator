package handler

import (
	"github.com/zdq0394/consul-operator/pkg/apis/consul/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func generateRedisConfigMap(rc *v1alpha1.Consul,
	labels map[string]string, ownerRefs []metav1.OwnerReference) *corev1.ConfigMap {
	name := generateName(configMapNamePrefix, rc.Name)
	namespace := rc.Namespace
	conf := `appendonly yes
cluster-enabled yes
cluster-config-file /var/lib/redis/nodes.conf
cluster-node-timeout 5000
dir /var/lib/redis
port 6379`
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:            name,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: ownerRefs,
		},
		Data: map[string]string{
			"redis.conf": conf,
		},
	}
}

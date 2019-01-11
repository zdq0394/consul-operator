package handler

import (
	"fmt"

	"encoding/json"

	"github.com/zdq0394/consul-operator/pkg/apis/consul/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type dnsConfig struct {
	EnableTruncate bool `json:"enable_truncate"`
	OnlyPassing    bool `json:"only_passing"`
}

type consulConf struct {
	BootstrapExpect    int       `json:"bootstrap_expect"`
	ClientAddr         string    `json:"client_addr"`
	DataCenter         string    `json:"datacenter"`
	DataDir            string    `json:"data_dir"`
	Domain             string    `json:"domain"`
	EnableScriptChecks bool      `json:"enable_script_checks"`
	DNSConfig          dnsConfig `json:"dns_config"`
	EnableSyslog       bool      `json:"enable_syslog"`
	LeaveOnTerminate   bool      `json:"leave_on_terminate"`
	LogLevel           string    `json:"log_level"`
	RejoinAfterLeave   bool      `json:"rejoin_after_leave"`
	Server             bool      `json:"server"`
	StartJoin          []string  `json:"start_join"`
	UIEnabled          bool      `json:"ui"`
}

func generateConfigMap(rc *v1alpha1.Consul,
	labels map[string]string, ownerRefs []metav1.OwnerReference, clusterDomain string) *corev1.ConfigMap {
	name := generateName(configMapNamePrefix, rc.Name)
	namespace := rc.Namespace
	statefulSetName := generateName(statefulsetNamePrefix, rc.Name)
	svcName := generateName(svcHeadlessNamePrefix, rc.Name)
	postix := fmt.Sprintf(".%s.%s.svc.k8s.%s", svcName, namespace, clusterDomain)
	conf := consulConf{
		BootstrapExpect:    int(rc.Spec.Consul.Replicas),
		ClientAddr:         "0.0.0.0",
		DataCenter:         "CN-East",
		DataDir:            "/var/consul",
		Domain:             "consul",
		EnableScriptChecks: true,
		DNSConfig: dnsConfig{
			EnableTruncate: true,
			OnlyPassing:    true,
		},
		EnableSyslog:     false,
		LeaveOnTerminate: true,
		LogLevel:         "INFO",
		RejoinAfterLeave: true,
		Server:           true,
		StartJoin: []string{
			fmt.Sprintf("%s-0.%s", statefulSetName, postix),
			fmt.Sprintf("%s-1.%s", statefulSetName, postix),
			fmt.Sprintf("%s-2.%s", statefulSetName, postix),
		},
		UIEnabled: true,
	}

	confData, err := json.Marshal(conf)
	if err != nil {
		return nil
	}
	confText := string(confData)

	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:            name,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: ownerRefs,
		},
		Data: map[string]string{
			"consul.json": confText,
		},
	}
}

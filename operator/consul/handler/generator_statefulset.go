package handler

import (
	"github.com/zdq0394/consul-operator/pkg/apis/consul/v1alpha1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func generateConsulStatefulset(rc *v1alpha1.Consul,
	labels map[string]string, ownerRefs []metav1.OwnerReference) *appsv1beta2.StatefulSet {
	name := generateName(statefulsetNamePrefix, rc.Name)
	serviceName := generateName(svcHeadlessNamePrefix, rc.Name)
	namespace := rc.Namespace

	spec := rc.Spec
	consulImage := spec.Consul.Image
	replicas := spec.Consul.Replicas

	ss := &appsv1beta2.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:            name,
			Namespace:       namespace,
			Labels:          labels,
			OwnerReferences: ownerRefs,
		},
		Spec: appsv1beta2.StatefulSetSpec{
			ServiceName: serviceName,
			Replicas:    &replicas,
			UpdateStrategy: appsv1beta2.StatefulSetUpdateStrategy{
				Type: "RollingUpdate",
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: &terminationGracePeriodSeconds,
					Containers: []corev1.Container{
						{
							Name:            "consul",
							Image:           consulImage,
							ImagePullPolicy: "Always",
							Args: []string{
								"agent",
								"-config-dir=/etc/consul.d",
							},
							Ports:        getContainerPorts(rc),
							VolumeMounts: getVolumeMounts(rc),
							Resources:    getResources(rc),
						},
					},
					Volumes: getVolumes(rc),
				},
			},
			VolumeClaimTemplates: getVolumeClaimTemplates(rc, labels, ownerRefs),
		},
	}
	return ss
}

func getVolumeMounts(rc *v1alpha1.Consul) []corev1.VolumeMount {
	return []corev1.VolumeMount{
		{
			Name:      "consul-conf",
			MountPath: "/etc/consul.d",
		},
		{
			Name:      "consul-data",
			MountPath: "/var/consul",
		},
	}
}

func getVolumes(rc *v1alpha1.Consul) []corev1.Volume {
	return []corev1.Volume{
		{
			Name: "consul-conf",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: generateName(configMapNamePrefix, rc.Name),
					},
				},
			},
		},
	}
}

func getContainerPorts(rc *v1alpha1.Consul) []corev1.ContainerPort {
	return []corev1.ContainerPort{
		{
			Name:          consulHeadlessPortName,
			ContainerPort: consulHeadlessPort,
			Protocol:      corev1.ProtocolTCP,
		},
	}
}

func getVolumeClaimTemplates(rc *v1alpha1.Consul,
	labels map[string]string, ownerRefs []metav1.OwnerReference) []corev1.PersistentVolumeClaim {
	storageSize := rc.Spec.Consul.Storage.Size
	storageClassName := rc.Spec.Consul.Storage.StorageClassName
	return []corev1.PersistentVolumeClaim{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:            "consul-data",
				Namespace:       rc.Namespace,
				Labels:          labels,
				OwnerReferences: ownerRefs,
			},
			Spec: corev1.PersistentVolumeClaimSpec{
				AccessModes: []corev1.PersistentVolumeAccessMode{
					"ReadWriteOnce",
				},
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: resource.MustParse(storageSize),
					},
				},
				StorageClassName: &storageClassName,
			},
		},
	}
}

func generateResourceList(cpu string, memory string) corev1.ResourceList {
	resources := corev1.ResourceList{}
	if cpu != "" {
		resources[corev1.ResourceCPU], _ = resource.ParseQuantity(cpu)
	}
	if memory != "" {
		resources[corev1.ResourceMemory], _ = resource.ParseQuantity(memory)
	}
	return resources
}

func getRequests(resources v1alpha1.Resources) corev1.ResourceList {
	return generateResourceList(resources.Requests.CPU, resources.Requests.Memory)
}

func getLimits(resources v1alpha1.Resources) corev1.ResourceList {
	return generateResourceList(resources.Limits.CPU, resources.Limits.Memory)
}

func getResources(r *v1alpha1.Consul) corev1.ResourceRequirements {
	return corev1.ResourceRequirements{
		Requests: getRequests(r.Spec.Consul.Resources),
		Limits:   getLimits(r.Spec.Consul.Resources),
	}
}

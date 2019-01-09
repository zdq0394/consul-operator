/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Consul is a specification for a Consul resource
type Consul struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConsulSpec   `json:"spec"`
	Status ConsulStatus `json:"status"`
}

// ConsulSpec is the spec for a ConsulSpec resource
type ConsulSpec struct {
	Consul ConsulSetting `json:"consul"`
}

// ConsulSetting settings of consul
type ConsulSetting struct {
	Replicas  int32       `json:"replicas,omitempty"`
	Resources Resources   `json:"resources,omitempty"`
	Image     string      `json:"image,omitempty"`
	Storage   StorageSpec `json:"storage,omitempty"`
}

// ConsulStatus is the status for a Consul resource
type ConsulStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConsulList is a list of Consul resources
type ConsulList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Consul `json:"items"`
}

// StorageSpec spec of storage.
type StorageSpec struct {
	Size             string `json:"size,omitempty"`
	StorageClassName string `json:"storageClassName,omitempty"`
}

// Resources sets the limits and requests for a container
type Resources struct {
	Requests CPUAndMem `json:"requests,omitempty"`
	Limits   CPUAndMem `json:"limits,omitempty"`
}

// CPUAndMem defines how many cpu and ram the container will request/limit
type CPUAndMem struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

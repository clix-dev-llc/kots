/*
Copyright 2019 Replicated, Inc..

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KurlValuesSpec defines the desired state of KurlValuesSpec
type KurlValuesSpec struct {
	IsAirgapped bool `json:"isAirgapped,omitempty"`
	IsKurl bool `json:"isKurl,omitempty"`
	IsKurlHA bool `json:"isKurlHA,omitempty"`
	LoadBalancerAddress string `json:"LoadBalancerAddress,omitempty"`
}

// KurlValuesStatus defines the observed state of KurlValues
type KurlValuesStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KurlValues is the Schema for the airgap API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type KurlValues struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KurlValuesSpec   `json:"spec,omitempty"`
	Status KurlValuesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KurlValuesList contains a list of KurlValuess
type KurlValuesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KurlValues `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KurlValues{}, &KurlValuesList{})
}
/*
Copyright 2020 hurz.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type EtcdRestorePhase string

var (
	EtcdRestoreRunning   EtcdDumpPhase = "Running"
	EtcdRestoreCompleted EtcdDumpPhase = "Completed"
	EtcdRestoreFailed    EtcdDumpPhase = "Failed"
)

// EtcdRestoreSpec defines the desired state of EtcdRestore
type EtcdRestoreSpec struct {
	ClusterReference string          `json:"cluster_reference"`
	Location         string          `json:"location"`
	Storage          StorageProvider `json:"storage"`
}

// EtcdRestoreStatus defines the observed state of EtcdRestore
type EtcdRestoreStatus struct {
	Conditions []EtcdRestoreCondition `json:"conditions"`
	Phase      EtcdRestorePhase       `json:"phase"`
}

type EtcdRestoreCondition struct {
	Ready                 bool        `json:"ready"`
	Reason                string      `json:"reason,omitempty"`
	Message               string      `json:"message,omitempty"`
	Location              string      `json:"location,omitempty"`
	LastedTranslationTime metav1.Time `json:"lasted_translation_time"`
}

// +kubebuilder:object:root=true

// EtcdRestore is the Schema for the etcdrestores API
type EtcdRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdRestoreSpec   `json:"spec,omitempty"`
	Status EtcdRestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EtcdRestoreList contains a list of EtcdRestore
type EtcdRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EtcdRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EtcdRestore{}, &EtcdRestoreList{})
}

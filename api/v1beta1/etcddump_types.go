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

type EtcdDumpPhase string

var (
	EtcdDumpRunning   EtcdDumpPhase = "Running"
	EtcdDumpCompleted EtcdDumpPhase = "Completed"
	EtcdDumpFailed    EtcdDumpPhase = "Failed"
)

// EtcdDumpSpec defines the desired state of EtcdDump
type EtcdDumpSpec struct {
	Scheduler        string          `json:"scheduler,omitempty"`
	ClusterReference string          `json:"cluster_reference"`
	Storage          StorageProvider `json:"storage"`
}

// EtcdDumpStatus defines the observed state of EtcdDump
type EtcdDumpStatus struct {
	Conditions []EtcdDumpCondition `json:"conditions"`
	Phase      EtcdDumpPhase       `json:"phase"`
}

type EtcdDumpCondition struct {
	Ready                 bool        `json:"ready"`
	Reason                string      `json:"reason,omitempty"`
	Message               string      `json:"message,omitempty"`
	Location              string      `json:"location,omitempty"`
	LastedTranslationTime metav1.Time `json:"lasted_translation_time"`
}

type StorageProvider struct {
	Tencent *TencentOSS `json:"tencent,ommitempt"`
}

type TencentOSS struct {
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	Url       string `json:"url"`
	Bucket    string `json:"bucket"`
}

// +kubebuilder:object:root=true

// EtcdDump is the Schema for the etcddumps API
type EtcdDump struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdDumpSpec   `json:"spec,omitempty"`
	Status EtcdDumpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EtcdDumpList contains a list of EtcdDump
type EtcdDumpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EtcdDump `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EtcdDump{}, &EtcdDumpList{})
}

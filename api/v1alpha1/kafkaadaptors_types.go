/*
Copyright 2025 ShonPhand.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type KafkaConnection struct {
	BootstrapServer []string `json:"bootstrapServer"`
	Host            string   `json:"host"`
	Username        string   `json:"username"`
	Password        string   `json:"password"`
	CACert          string   `json:"caCert"`
}

type Transformation struct {
	Name               string `json:"name"`
	SourceTopic        string `json:"sourceTopic"`
	TargetTopic        string `json:"targetTopic"`
	TransformationType string `json:"transformationType"`
}

// KafkaAdaptorsSpec defines the desired state of KafkaAdaptors.
type KafkaAdaptorsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of KafkaAdaptors. Edit kafkaadaptors_types.go to remove/update
	KafkaConnection KafkaConnection  `json:"kafkaConnection"`
	Transformations []Transformation `json:"transformations"`
	Image           string           `json:"image"`
}

// KafkaAdaptorsStatus defines the observed state of KafkaAdaptors.
type KafkaAdaptorsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KafkaAdaptors is the Schema for the kafkaadaptors API.
type KafkaAdaptors struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaAdaptorsSpec   `json:"spec,omitempty"`
	Status KafkaAdaptorsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaAdaptorsList contains a list of KafkaAdaptors.
type KafkaAdaptorsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaAdaptors `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KafkaAdaptors{}, &KafkaAdaptorsList{})
}

/*
Copyright 2025.

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

type KymaPolicyGroup struct {
	Name string `json:"name,omitempty"`
	// The list of kyvernoPolicies that belong to this group
	KyvernoPolicies []string `json:"kyvernoPolicies,omitempty"`
	DefaultEnabled  bool     `json:"defaultEnabled,omitempty"`
	Enabled         bool     `json:"enabled,omitempty"`
}

// KymaPolicySpec defines the desired state of KymaPolicy.
type KymaPolicyConfigSpec struct {
	// List of KymaPolicyGroups
	PolicyGroups  []KymaPolicyGroup `json:"items,omitempty"`
	DefaultPolicy string            `json:"defaultPolicy"`

	// In intrusive mode, Kyverno blocks policy violating actions. In non-intrusive Kyverno only logs violating actions
	IntrusiveMode bool `json:"intrusiveMode,omitempty"`
}

// KymaPolicyStatus defines the observed state of KymaPolicy.
type KymaPolicyConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KymaPolicy is the Schema for the kymapolicies API.
type KymaPolicyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KymaPolicyConfigSpec   `json:"spec,omitempty"`
	Status KymaPolicyConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KymaPolicyList contains a list of KymaPolicy.
type KymaPolicyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KymaPolicyConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KymaPolicyConfig{}, &KymaPolicyConfigList{})
}

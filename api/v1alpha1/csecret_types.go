/*
Copyright 2024.

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

type AuthSecretRef struct {
    Name string `json:"name"`
    Key  string `json:"key,omitempty"` // Optional, defaults to "credentials.json"
}

// CsecretSpec defines the desired state of Csecret
type CsecretSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Csecret. Edit csecret_types.go to remove/update
	SecretName string `json:"secretName"`

	ProjectID string `json:"projectId"`

	CheckSecretsSeconds *int64 `json:"CheckSecretsSeconds"`

	SecretRef AuthSecretRef `json:"SecretRef"`

}

// CsecretStatus defines the observed state of Csecret
type CsecretStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Csecret is the Schema for the csecrets API
type Csecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CsecretSpec   `json:"spec,omitempty"`
	Status CsecretStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CsecretList contains a list of Csecret
type CsecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Csecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Csecret{}, &CsecretList{})
}

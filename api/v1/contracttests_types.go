/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ContractTestsSpec defines the desired state of ContractTests
type ContractTestsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ContractTests. Edit contracttests_types.go to remove/update
	ContractName string `json:"contractName,omitempty"`
	WaitSecs     int    `json:"waitSeconds,omitempty"`
}

// ContractTestsStatus defines the observed state of ContractTests
type ContractTestsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	UpdatedAt string `json:"updatedAt,omitempty"`
	Message   string `json:"message,omitempty"`
	Status    string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ContractTests is the Schema for the contracttests API
type ContractTests struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContractTestsSpec   `json:"spec,omitempty"`
	Status ContractTestsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ContractTestsList contains a list of ContractTests
type ContractTestsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContractTests `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContractTests{}, &ContractTestsList{})
}

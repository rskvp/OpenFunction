/*


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

const (
	Knative = "Knative"
)

// ServingSpec defines the desired state of Serving
type ServingSpec struct {
	// Function image name
	Image string `json:"image"`
	// The port on which the function will be invoked
	Port *string `json:"port,omitempty"`
	// The backend runtime to run a function, for example Knative
	Runtime *string `json:"runtime,omitempty"`
}

// ServingStatus defines the observed state of Serving
type ServingStatus struct {
	Phase string `json:"phase,omitempty"`
	State string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=fs
// +kubebuilder:subresource:status
// Serving is the Schema for the servings API
type Serving struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServingSpec   `json:"spec,omitempty"`
	Status ServingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServingList contains a list of Serving
type ServingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Serving `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Serving{}, &ServingList{})
}
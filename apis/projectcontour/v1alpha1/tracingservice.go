// Copyright Project Contour Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	contour_api_v1 "github.com/projectcontour/contour/apis/projectcontour/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TracingServiceSpec defines the desired state of a TracingService resource.
type TracingServiceSpec struct {
	// Format is the type of Tracing configured in Envoy.
	// Currently, only opencensus is supported.
	// +kubebuilder:validation:Enum=opencensus
	Format *string `json:"format,omitempty"`

	// AgentAddress is the address for the Collector
	// to report traces to. The specific type is
	AgentAddress *string `json:"address,omitempty"`
}

// TracingServiceStatus defines the observed state of an
// TracingService resource.
type TracingServiceStatus struct {
	// Conditions contains the current status of the TracingService resource.
	//
	// Contour will update a single condition, `Valid`, that is in normal-true polarity.
	//
	// Contour will not modify any other Conditions set in this block,
	// in case some other controller wants to add a Condition.
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []contour_api_v1.DetailedCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=tracingservice;tracingservices

// TracingService is the schema for the Contour Tracing Service API.
// A TracingService resource binds an OpenCensus Cluster to Envoy
// in the bootstrap phase so that HTTPProxy Resources can specify tracing
type TracingService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TracingServiceSpec   `json:"spec,omitempty"`
	Status TracingServiceStatus `json:"status,omitempty"`
}

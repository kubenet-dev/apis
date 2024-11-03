/*
Copyright 2024 Nokia.

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
	"reflect"

	condv1alpha1 "github.com/kform-dev/choreo/apis/condition/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NodeTemplateSpec defines the desired state of NodeTemplate
type NodeTemplateSpec struct {
	Provider     string              `json:"provider" protobuf:"bytes,1,opt,name=provider"`
	PlatformType string              `json:"platformType" protobuf:"bytes,2,opt,name=platformType"`
	Ports        []*NodeTemplatePort `json:"ports" protobuf:"bytes,2,opt,name=ports"`
}

type NodeTemplatePort struct {
	Ids     NodeTemplatePortIds     `json:"ids" protobuf:"bytes,1,opt,name=ids"`
	Adaptor NodeTemplatePortAdaptor `json:"adaptor" protobuf:"bytes,2,opt,name=adaptor"`
}

type NodeTemplatePortIds struct {
	Start uint `json:"start" protobuf:"bytes,1,opt,name=start"`
	End  uint `json:"end" protobuf:"bytes,2,opt,name=end"`
}

type NodeTemplatePortAdaptor struct {
	Name       string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Speed      string `json:"speed" protobuf:"bytes,2,opt,name=speed"`
	Pluggable  bool   `json:"pluggable" protobuf:"bytes,3,opt,name=pluggable"`
	Connectors uint   `json:"connectors" protobuf:"bytes,4,opt,name=connectors"`
}

// NodeTemplateStatus defines the observed state of NodeTemplate
type NodeTemplateStatus struct {
	// ConditionedStatus provides the status of the Readiness using conditions
	// if the condition is true the other attributes in the status are meaningful
	condv1alpha1.ConditionedStatus `json:",inline" protobuf:"bytes,1,opt,name=conditionedStatus"`
}

// +genclient
// +k8s:deepcopy-gen:NodeTemplates=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:resource:categories={kubenet}

// NodeTemplate defines the Schema for the NodeTemplate API
type NodeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   NodeTemplateSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status NodeTemplateStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// NodeTemplateList contains a list of NodeTemplates
// +k8s:deepcopy-gen:NodeTemplates=k8s.io/apimachinery/pkg/runtime.Object
type NodeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []NodeTemplate `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// NodeTemplate type metadata.
var (
	NodeTemplateKind = reflect.TypeOf(NodeTemplate{}).Name()
)

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

package v1alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cat
// +k8s:openapi-gen=true
type Cat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CatSpec   `json:"spec,omitempty"`
	Status CatStatus `json:"status,omitempty"`
}

// CatList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Cat `json:"items"`
}

// CatSpec defines the desired state of Cat
type CatSpec struct {
	Name string `json:"name"`
}

var _ resource.Object = &Cat{}
var _ resourcestrategy.Validater = &Cat{}

func (in *Cat) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *Cat) NamespaceScoped() bool {
	return false
}

func (in *Cat) New() runtime.Object {
	return &Cat{}
}

func (in *Cat) NewList() runtime.Object {
	return &CatList{}
}

func (in *Cat) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "animal.hm.com",
		Version:  "v1alpha1",
		Resource: "cats",
	}
}

func (in *Cat) IsStorageVersion() bool {
	return true
}

func (in *Cat) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	allErrs := field.ErrorList{}
	if len(in.Spec.Name) == 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "name"), in.Spec.Name, "must be specify"))
	}
	return allErrs
}

var _ resource.ObjectList = &CatList{}

func (in *CatList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// CatStatus defines the observed state of Cat
type CatStatus struct {
}

func (in CatStatus) SubResourceName() string {
	return "status"
}

// Cat implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &Cat{}

func (in *Cat) GetStatus() resource.StatusSubResource {
	return in.Status
}

// CatStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &CatStatus{}

func (in CatStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*Cat).Status = in
}

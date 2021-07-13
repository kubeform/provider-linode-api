/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Stackscript struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackscriptSpec   `json:"spec,omitempty"`
	Status            StackscriptStatus `json:"status,omitempty"`
}

type StackscriptSpecUserDefinedFields struct {
	// The default value. If not specified, this value will be used.
	// +optional
	Default *string `json:"default,omitempty" tf:"default"`
	// An example value for the field.
	// +optional
	Example *string `json:"example,omitempty" tf:"example"`
	// A human-readable label for the field that will serve as the input prompt for entering the value during deployment.
	// +optional
	Label *string `json:"label,omitempty" tf:"label"`
	// A list of acceptable values for the field in any quantity, combination or order.
	// +optional
	ManyOf *string `json:"manyOf,omitempty" tf:"many_of"`
	// The name of the field.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// A list of acceptable single values for the field.
	// +optional
	OneOf *string `json:"oneOf,omitempty" tf:"one_of"`
}

type StackscriptSpec struct {
	State *StackscriptSpecResource `json:"state,omitempty" tf:"-"`

	Resource StackscriptSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type StackscriptSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The date this StackScript was created.
	// +optional
	Created *string `json:"created,omitempty" tf:"created"`
	// Count of currently active, deployed Linodes created from this StackScript.
	// +optional
	DeploymentsActive *int64 `json:"deploymentsActive,omitempty" tf:"deployments_active"`
	// The total number of times this StackScript has been deployed.
	// +optional
	DeploymentsTotal *int64 `json:"deploymentsTotal,omitempty" tf:"deployments_total"`
	// A description for the StackScript.
	Description *string `json:"description" tf:"description"`
	// An array of Image IDs representing the Images that this StackScript is compatible for deploying with.
	Images []string `json:"images" tf:"images"`
	// This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.
	// +optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public"`
	// The StackScript's label is for display purposes only.
	Label *string `json:"label" tf:"label"`
	// This field allows you to add notes for the set of revisions made to this StackScript.
	// +optional
	RevNote *string `json:"revNote,omitempty" tf:"rev_note"`
	// The script to execute when provisioning a new Linode with this StackScript.
	Script *string `json:"script" tf:"script"`
	// The date this StackScript was updated.
	// +optional
	Updated *string `json:"updated,omitempty" tf:"updated"`
	// This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.
	// +optional
	UserDefinedFields []StackscriptSpecUserDefinedFields `json:"userDefinedFields,omitempty" tf:"user_defined_fields"`
	// The Gravatar ID for the User who created the StackScript.
	// +optional
	UserGravatarID *string `json:"userGravatarID,omitempty" tf:"user_gravatar_id"`
	// The User who created the StackScript.
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type StackscriptStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StackscriptList is a list of Stackscripts
type StackscriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Stackscript CRD objects
	Items []Stackscript `json:"items,omitempty"`
}

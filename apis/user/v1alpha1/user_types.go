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

type User struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec,omitempty"`
	Status            UserStatus `json:"status,omitempty"`
}

type UserSpecDomainGrant struct {
	// The ID of the entity this grant applies to.
	ID *int64 `json:"ID" tf:"id"`
	// The level of access this User has to this entity. If null, this User has no access.
	Permissions *string `json:"permissions" tf:"permissions"`
}

type UserSpecGlobalGrants struct {
	// The level of access this User has to Account-level actions, like billing information. A restricted User will never be able to manage users.
	// +optional
	AccountAccess *string `json:"accountAccess,omitempty" tf:"account_access"`
	// If true, this User may add Domains.
	// +optional
	AddDomains *bool `json:"addDomains,omitempty" tf:"add_domains"`
	// If true, this User may add Images.
	// +optional
	AddImages *bool `json:"addImages,omitempty" tf:"add_images"`
	// If true, this User may create Linodes.
	// +optional
	AddLinodes *bool `json:"addLinodes,omitempty" tf:"add_linodes"`
	// If true, this User may create Longview clients and view the current plan.
	// +optional
	AddLongview *bool `json:"addLongview,omitempty" tf:"add_longview"`
	// If true, this User may add NodeBalancers.
	// +optional
	AddNodebalancers *bool `json:"addNodebalancers,omitempty" tf:"add_nodebalancers"`
	// If true, this User may add StackScripts.
	// +optional
	AddStackscripts *bool `json:"addStackscripts,omitempty" tf:"add_stackscripts"`
	// If true, this User may add Volumes.
	// +optional
	AddVolumes *bool `json:"addVolumes,omitempty" tf:"add_volumes"`
	// If true, this User may cancel the entire Account.
	// +optional
	CancelAccount *bool `json:"cancelAccount,omitempty" tf:"cancel_account"`
	// If true, this User may manage the Account’s Longview subscription.
	// +optional
	LongviewSubscription *bool `json:"longviewSubscription,omitempty" tf:"longview_subscription"`
}

type UserSpecImageGrant struct {
	// The ID of the entity this grant applies to.
	ID *int64 `json:"ID" tf:"id"`
	// The level of access this User has to this entity. If null, this User has no access.
	Permissions *string `json:"permissions" tf:"permissions"`
}

type UserSpecLinodeGrant struct {
	// The ID of the entity this grant applies to.
	ID *int64 `json:"ID" tf:"id"`
	// The level of access this User has to this entity. If null, this User has no access.
	Permissions *string `json:"permissions" tf:"permissions"`
}

type UserSpecLongviewGrant struct {
	// The ID of the entity this grant applies to.
	ID *int64 `json:"ID" tf:"id"`
	// The level of access this User has to this entity. If null, this User has no access.
	Permissions *string `json:"permissions" tf:"permissions"`
}

type UserSpecNodebalancerGrant struct {
	// The ID of the entity this grant applies to.
	ID *int64 `json:"ID" tf:"id"`
	// The level of access this User has to this entity. If null, this User has no access.
	Permissions *string `json:"permissions" tf:"permissions"`
}

type UserSpecStackscriptGrant struct {
	// The ID of the entity this grant applies to.
	ID *int64 `json:"ID" tf:"id"`
	// The level of access this User has to this entity. If null, this User has no access.
	Permissions *string `json:"permissions" tf:"permissions"`
}

type UserSpecVolumeGrant struct {
	// The ID of the entity this grant applies to.
	ID *int64 `json:"ID" tf:"id"`
	// The level of access this User has to this entity. If null, this User has no access.
	Permissions *string `json:"permissions" tf:"permissions"`
}

type UserSpec struct {
	State *UserSpecResource `json:"state,omitempty" tf:"-"`

	Resource UserSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type UserSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A set containing all of the user's active grants.
	// +optional
	DomainGrant []UserSpecDomainGrant `json:"domainGrant,omitempty" tf:"domain_grant"`
	// The email of the user.
	Email *string `json:"email" tf:"email"`
	// A structure containing the Account-level grants a User has.
	// +optional
	GlobalGrants *UserSpecGlobalGrants `json:"globalGrants,omitempty" tf:"global_grants"`
	// A set containing all of the user's active grants.
	// +optional
	ImageGrant []UserSpecImageGrant `json:"imageGrant,omitempty" tf:"image_grant"`
	// A set containing all of the user's active grants.
	// +optional
	LinodeGrant []UserSpecLinodeGrant `json:"linodeGrant,omitempty" tf:"linode_grant"`
	// A set containing all of the user's active grants.
	// +optional
	LongviewGrant []UserSpecLongviewGrant `json:"longviewGrant,omitempty" tf:"longview_grant"`
	// A set containing all of the user's active grants.
	// +optional
	NodebalancerGrant []UserSpecNodebalancerGrant `json:"nodebalancerGrant,omitempty" tf:"nodebalancer_grant"`
	// If true, the user must be explicitly granted access to platform actions and entities.
	// +optional
	Restricted *bool `json:"restricted,omitempty" tf:"restricted"`
	// SSH keys to add to the user profile.
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`
	// A set containing all of the user's active grants.
	// +optional
	StackscriptGrant []UserSpecStackscriptGrant `json:"stackscriptGrant,omitempty" tf:"stackscript_grant"`
	// If the User has Two Factor Authentication (TFA) enabled.
	// +optional
	TfaEnabled *bool `json:"tfaEnabled,omitempty" tf:"tfa_enabled"`
	// The username of the user.
	Username *string `json:"username" tf:"username"`
	// A set containing all of the user's active grants.
	// +optional
	VolumeGrant []UserSpecVolumeGrant `json:"volumeGrant,omitempty" tf:"volume_grant"`
}

type UserStatus struct {
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

// UserList is a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of User CRD objects
	Items []User `json:"items,omitempty"`
}

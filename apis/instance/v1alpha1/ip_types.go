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

type Ip struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpSpec   `json:"spec,omitempty"`
	Status            IpStatus `json:"status,omitempty"`
}

type IpSpec struct {
	State *IpSpecResource `json:"state,omitempty" tf:"-"`

	Resource IpSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type IpSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The resulting IPv4 address.
	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// If true, the instance will be rebooted to update network interfaces.
	// +optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`
	// The default gateway for this address
	// +optional
	Gateway *string `json:"gateway,omitempty" tf:"gateway"`
	// The ID of the Linode to allocate an IPv4 address for.
	LinodeID *int64 `json:"linodeID" tf:"linode_id"`
	// The number of bits set in the subnet mask.
	// +optional
	Prefix *int64 `json:"prefix,omitempty" tf:"prefix"`
	// Whether the IPv4 address is public or private.
	// +optional
	Public *bool `json:"public,omitempty" tf:"public"`
	// The reverse DNS assigned to this address.
	// +optional
	Rdns *string `json:"rdns,omitempty" tf:"rdns"`
	// The region this IP resides in.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The mask that separates host bits from network bits for this address.
	// +optional
	SubnetMask *string `json:"subnetMask,omitempty" tf:"subnet_mask"`
	// The type of IP address.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type IpStatus struct {
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

// IpList is a list of Ips
type IpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ip CRD objects
	Items []Ip `json:"items,omitempty"`
}

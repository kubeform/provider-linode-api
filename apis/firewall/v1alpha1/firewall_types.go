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

type Firewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec,omitempty"`
	Status            FirewallStatus `json:"status,omitempty"`
}

type FirewallSpecDevices struct {
	// The ID of the underlying entity for the firewall device (e.g. the Linode's ID).
	// +optional
	EntityID *int64 `json:"entityID,omitempty" tf:"entity_id"`
	// The ID of the firewall device.
	// +optional
	ID *int64 `json:"ID,omitempty" tf:"id"`
	// The label of the underlying entity for the firewall device.
	// +optional
	Label *string `json:"label,omitempty" tf:"label"`
	// The type of firewall device.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The URL of the underlying entity for the firewall device.
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type FirewallSpecInbound struct {
	// Controls whether traffic is accepted or dropped by this rule. Overrides the Firewall???s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.
	Action *string `json:"action" tf:"action"`
	// A list of IP addresses, CIDR blocks, or 0.0.0.0/0 (to allow all) this rule applies to.
	// +optional
	Ipv4 []string `json:"ipv4,omitempty" tf:"ipv4"`
	// A list of IPv6 addresses or networks this rule applies to.
	// +optional
	Ipv6 []string `json:"ipv6,omitempty" tf:"ipv6"`
	// Used to identify this rule. For display purposes only.
	Label *string `json:"label" tf:"label"`
	// A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").
	// +optional
	Ports *string `json:"ports,omitempty" tf:"ports"`
	// The network protocol this rule controls.
	Protocol *string `json:"protocol" tf:"protocol"`
}

type FirewallSpecOutbound struct {
	// Controls whether traffic is accepted or dropped by this rule. Overrides the Firewall???s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.
	Action *string `json:"action" tf:"action"`
	// A list of IP addresses, CIDR blocks, or 0.0.0.0/0 (to allow all) this rule applies to.
	// +optional
	Ipv4 []string `json:"ipv4,omitempty" tf:"ipv4"`
	// A list of IPv6 addresses or networks this rule applies to.
	// +optional
	Ipv6 []string `json:"ipv6,omitempty" tf:"ipv6"`
	// Used to identify this rule. For display purposes only.
	Label *string `json:"label" tf:"label"`
	// A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").
	// +optional
	Ports *string `json:"ports,omitempty" tf:"ports"`
	// The network protocol this rule controls.
	Protocol *string `json:"protocol" tf:"protocol"`
}

type FirewallSpec struct {
	State *FirewallSpecResource `json:"state,omitempty" tf:"-"`

	Resource FirewallSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FirewallSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The devices associated with this firewall.
	// +optional
	Devices []FirewallSpecDevices `json:"devices,omitempty" tf:"devices"`
	// If true, the Firewall is inactive.
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// A firewall rule that specifies what inbound network traffic is allowed.
	// +optional
	Inbound []FirewallSpecInbound `json:"inbound,omitempty" tf:"inbound"`
	// The default behavior for inbound traffic. This setting can be overridden by updating the inbound.action property for an individual Firewall Rule.
	InboundPolicy *string `json:"inboundPolicy" tf:"inbound_policy"`
	// The label for the Firewall. For display purposes only. If no label is provided, a default will be assigned.
	Label *string `json:"label" tf:"label"`
	// The IDs of Linodes to apply this firewall to.
	// +optional
	Linodes []int64 `json:"linodes,omitempty" tf:"linodes"`
	// A firewall rule that specifies what outbound network traffic is allowed.
	// +optional
	Outbound []FirewallSpecOutbound `json:"outbound,omitempty" tf:"outbound"`
	// The default behavior for outbound traffic. This setting can be overridden by updating the outbound.action property for an individual Firewall Rule.
	OutboundPolicy *string `json:"outboundPolicy" tf:"outbound_policy"`
	// The status of the firewall.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
}

type FirewallStatus struct {
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

// FirewallList is a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Firewall CRD objects
	Items []Firewall `json:"items,omitempty"`
}

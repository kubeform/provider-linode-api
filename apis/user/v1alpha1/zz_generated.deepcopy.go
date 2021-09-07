// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserList) DeepCopyInto(out *UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserList.
func (in *UserList) DeepCopy() *UserList {
	if in == nil {
		return nil
	}
	out := new(UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(UserSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecDomainGrant) DeepCopyInto(out *UserSpecDomainGrant) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecDomainGrant.
func (in *UserSpecDomainGrant) DeepCopy() *UserSpecDomainGrant {
	if in == nil {
		return nil
	}
	out := new(UserSpecDomainGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecGlobalGrants) DeepCopyInto(out *UserSpecGlobalGrants) {
	*out = *in
	if in.AccountAccess != nil {
		in, out := &in.AccountAccess, &out.AccountAccess
		*out = new(string)
		**out = **in
	}
	if in.AddDomains != nil {
		in, out := &in.AddDomains, &out.AddDomains
		*out = new(bool)
		**out = **in
	}
	if in.AddImages != nil {
		in, out := &in.AddImages, &out.AddImages
		*out = new(bool)
		**out = **in
	}
	if in.AddLinodes != nil {
		in, out := &in.AddLinodes, &out.AddLinodes
		*out = new(bool)
		**out = **in
	}
	if in.AddLongview != nil {
		in, out := &in.AddLongview, &out.AddLongview
		*out = new(bool)
		**out = **in
	}
	if in.AddNodebalancers != nil {
		in, out := &in.AddNodebalancers, &out.AddNodebalancers
		*out = new(bool)
		**out = **in
	}
	if in.AddStackscripts != nil {
		in, out := &in.AddStackscripts, &out.AddStackscripts
		*out = new(bool)
		**out = **in
	}
	if in.AddVolumes != nil {
		in, out := &in.AddVolumes, &out.AddVolumes
		*out = new(bool)
		**out = **in
	}
	if in.CancelAccount != nil {
		in, out := &in.CancelAccount, &out.CancelAccount
		*out = new(bool)
		**out = **in
	}
	if in.LongviewSubscription != nil {
		in, out := &in.LongviewSubscription, &out.LongviewSubscription
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecGlobalGrants.
func (in *UserSpecGlobalGrants) DeepCopy() *UserSpecGlobalGrants {
	if in == nil {
		return nil
	}
	out := new(UserSpecGlobalGrants)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecImageGrant) DeepCopyInto(out *UserSpecImageGrant) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecImageGrant.
func (in *UserSpecImageGrant) DeepCopy() *UserSpecImageGrant {
	if in == nil {
		return nil
	}
	out := new(UserSpecImageGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecLinodeGrant) DeepCopyInto(out *UserSpecLinodeGrant) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecLinodeGrant.
func (in *UserSpecLinodeGrant) DeepCopy() *UserSpecLinodeGrant {
	if in == nil {
		return nil
	}
	out := new(UserSpecLinodeGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecLongviewGrant) DeepCopyInto(out *UserSpecLongviewGrant) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecLongviewGrant.
func (in *UserSpecLongviewGrant) DeepCopy() *UserSpecLongviewGrant {
	if in == nil {
		return nil
	}
	out := new(UserSpecLongviewGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecNodebalancerGrant) DeepCopyInto(out *UserSpecNodebalancerGrant) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecNodebalancerGrant.
func (in *UserSpecNodebalancerGrant) DeepCopy() *UserSpecNodebalancerGrant {
	if in == nil {
		return nil
	}
	out := new(UserSpecNodebalancerGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecResource) DeepCopyInto(out *UserSpecResource) {
	*out = *in
	if in.DomainGrant != nil {
		in, out := &in.DomainGrant, &out.DomainGrant
		*out = make([]UserSpecDomainGrant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.GlobalGrants != nil {
		in, out := &in.GlobalGrants, &out.GlobalGrants
		*out = new(UserSpecGlobalGrants)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageGrant != nil {
		in, out := &in.ImageGrant, &out.ImageGrant
		*out = make([]UserSpecImageGrant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LinodeGrant != nil {
		in, out := &in.LinodeGrant, &out.LinodeGrant
		*out = make([]UserSpecLinodeGrant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LongviewGrant != nil {
		in, out := &in.LongviewGrant, &out.LongviewGrant
		*out = make([]UserSpecLongviewGrant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodebalancerGrant != nil {
		in, out := &in.NodebalancerGrant, &out.NodebalancerGrant
		*out = make([]UserSpecNodebalancerGrant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Restricted != nil {
		in, out := &in.Restricted, &out.Restricted
		*out = new(bool)
		**out = **in
	}
	if in.SshKeys != nil {
		in, out := &in.SshKeys, &out.SshKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StackscriptGrant != nil {
		in, out := &in.StackscriptGrant, &out.StackscriptGrant
		*out = make([]UserSpecStackscriptGrant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TfaEnabled != nil {
		in, out := &in.TfaEnabled, &out.TfaEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.VolumeGrant != nil {
		in, out := &in.VolumeGrant, &out.VolumeGrant
		*out = make([]UserSpecVolumeGrant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecResource.
func (in *UserSpecResource) DeepCopy() *UserSpecResource {
	if in == nil {
		return nil
	}
	out := new(UserSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecStackscriptGrant) DeepCopyInto(out *UserSpecStackscriptGrant) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecStackscriptGrant.
func (in *UserSpecStackscriptGrant) DeepCopy() *UserSpecStackscriptGrant {
	if in == nil {
		return nil
	}
	out := new(UserSpecStackscriptGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecVolumeGrant) DeepCopyInto(out *UserSpecVolumeGrant) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecVolumeGrant.
func (in *UserSpecVolumeGrant) DeepCopy() *UserSpecVolumeGrant {
	if in == nil {
		return nil
	}
	out := new(UserSpecVolumeGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}

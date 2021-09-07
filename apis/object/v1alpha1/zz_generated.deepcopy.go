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
func (in *StorageBucket) DeepCopyInto(out *StorageBucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucket.
func (in *StorageBucket) DeepCopy() *StorageBucket {
	if in == nil {
		return nil
	}
	out := new(StorageBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageBucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketList) DeepCopyInto(out *StorageBucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageBucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketList.
func (in *StorageBucketList) DeepCopy() *StorageBucketList {
	if in == nil {
		return nil
	}
	out := new(StorageBucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageBucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketSpec) DeepCopyInto(out *StorageBucketSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(StorageBucketSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketSpec.
func (in *StorageBucketSpec) DeepCopy() *StorageBucketSpec {
	if in == nil {
		return nil
	}
	out := new(StorageBucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketSpecCert) DeepCopyInto(out *StorageBucketSpecCert) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketSpecCert.
func (in *StorageBucketSpecCert) DeepCopy() *StorageBucketSpecCert {
	if in == nil {
		return nil
	}
	out := new(StorageBucketSpecCert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketSpecLifecycleRule) DeepCopyInto(out *StorageBucketSpecLifecycleRule) {
	*out = *in
	if in.AbortIncompleteMultipartUploadDays != nil {
		in, out := &in.AbortIncompleteMultipartUploadDays, &out.AbortIncompleteMultipartUploadDays
		*out = new(int64)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = new(StorageBucketSpecLifecycleRuleExpiration)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NoncurrentVersionExpiration != nil {
		in, out := &in.NoncurrentVersionExpiration, &out.NoncurrentVersionExpiration
		*out = new(StorageBucketSpecLifecycleRuleNoncurrentVersionExpiration)
		(*in).DeepCopyInto(*out)
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketSpecLifecycleRule.
func (in *StorageBucketSpecLifecycleRule) DeepCopy() *StorageBucketSpecLifecycleRule {
	if in == nil {
		return nil
	}
	out := new(StorageBucketSpecLifecycleRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketSpecLifecycleRuleExpiration) DeepCopyInto(out *StorageBucketSpecLifecycleRuleExpiration) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	if in.ExpiredObjectDeleteMarker != nil {
		in, out := &in.ExpiredObjectDeleteMarker, &out.ExpiredObjectDeleteMarker
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketSpecLifecycleRuleExpiration.
func (in *StorageBucketSpecLifecycleRuleExpiration) DeepCopy() *StorageBucketSpecLifecycleRuleExpiration {
	if in == nil {
		return nil
	}
	out := new(StorageBucketSpecLifecycleRuleExpiration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketSpecLifecycleRuleNoncurrentVersionExpiration) DeepCopyInto(out *StorageBucketSpecLifecycleRuleNoncurrentVersionExpiration) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketSpecLifecycleRuleNoncurrentVersionExpiration.
func (in *StorageBucketSpecLifecycleRuleNoncurrentVersionExpiration) DeepCopy() *StorageBucketSpecLifecycleRuleNoncurrentVersionExpiration {
	if in == nil {
		return nil
	}
	out := new(StorageBucketSpecLifecycleRuleNoncurrentVersionExpiration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketSpecResource) DeepCopyInto(out *StorageBucketSpecResource) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.Acl != nil {
		in, out := &in.Acl, &out.Acl
		*out = new(string)
		**out = **in
	}
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = new(StorageBucketSpecCert)
		(*in).DeepCopyInto(*out)
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.CorsEnabled != nil {
		in, out := &in.CorsEnabled, &out.CorsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.LifecycleRule != nil {
		in, out := &in.LifecycleRule, &out.LifecycleRule
		*out = make([]StorageBucketSpecLifecycleRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		*out = new(string)
		**out = **in
	}
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketSpecResource.
func (in *StorageBucketSpecResource) DeepCopy() *StorageBucketSpecResource {
	if in == nil {
		return nil
	}
	out := new(StorageBucketSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketStatus) DeepCopyInto(out *StorageBucketStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketStatus.
func (in *StorageBucketStatus) DeepCopy() *StorageBucketStatus {
	if in == nil {
		return nil
	}
	out := new(StorageBucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageKey) DeepCopyInto(out *StorageKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageKey.
func (in *StorageKey) DeepCopy() *StorageKey {
	if in == nil {
		return nil
	}
	out := new(StorageKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageKeyList) DeepCopyInto(out *StorageKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageKeyList.
func (in *StorageKeyList) DeepCopy() *StorageKeyList {
	if in == nil {
		return nil
	}
	out := new(StorageKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageKeySpec) DeepCopyInto(out *StorageKeySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(StorageKeySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageKeySpec.
func (in *StorageKeySpec) DeepCopy() *StorageKeySpec {
	if in == nil {
		return nil
	}
	out := new(StorageKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageKeySpecBucketAccess) DeepCopyInto(out *StorageKeySpecBucketAccess) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageKeySpecBucketAccess.
func (in *StorageKeySpecBucketAccess) DeepCopy() *StorageKeySpecBucketAccess {
	if in == nil {
		return nil
	}
	out := new(StorageKeySpecBucketAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageKeySpecResource) DeepCopyInto(out *StorageKeySpecResource) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.BucketAccess != nil {
		in, out := &in.BucketAccess, &out.BucketAccess
		*out = make([]StorageKeySpecBucketAccess, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Limited != nil {
		in, out := &in.Limited, &out.Limited
		*out = new(bool)
		**out = **in
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageKeySpecResource.
func (in *StorageKeySpecResource) DeepCopy() *StorageKeySpecResource {
	if in == nil {
		return nil
	}
	out := new(StorageKeySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageKeyStatus) DeepCopyInto(out *StorageKeyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageKeyStatus.
func (in *StorageKeyStatus) DeepCopy() *StorageKeyStatus {
	if in == nil {
		return nil
	}
	out := new(StorageKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObject) DeepCopyInto(out *StorageObject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObject.
func (in *StorageObject) DeepCopy() *StorageObject {
	if in == nil {
		return nil
	}
	out := new(StorageObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageObject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObjectList) DeepCopyInto(out *StorageObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObjectList.
func (in *StorageObjectList) DeepCopy() *StorageObjectList {
	if in == nil {
		return nil
	}
	out := new(StorageObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObjectSpec) DeepCopyInto(out *StorageObjectSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(StorageObjectSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObjectSpec.
func (in *StorageObjectSpec) DeepCopy() *StorageObjectSpec {
	if in == nil {
		return nil
	}
	out := new(StorageObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObjectSpecResource) DeepCopyInto(out *StorageObjectSpecResource) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.Acl != nil {
		in, out := &in.Acl, &out.Acl
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.CacheControl != nil {
		in, out := &in.CacheControl, &out.CacheControl
		*out = new(string)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.ContentBase64 != nil {
		in, out := &in.ContentBase64, &out.ContentBase64
		*out = new(string)
		**out = **in
	}
	if in.ContentDisposition != nil {
		in, out := &in.ContentDisposition, &out.ContentDisposition
		*out = new(string)
		**out = **in
	}
	if in.ContentEncoding != nil {
		in, out := &in.ContentEncoding, &out.ContentEncoding
		*out = new(string)
		**out = **in
	}
	if in.ContentLanguage != nil {
		in, out := &in.ContentLanguage, &out.ContentLanguage
		*out = new(string)
		**out = **in
	}
	if in.ContentType != nil {
		in, out := &in.ContentType, &out.ContentType
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.VersionID != nil {
		in, out := &in.VersionID, &out.VersionID
		*out = new(string)
		**out = **in
	}
	if in.WebsiteRedirect != nil {
		in, out := &in.WebsiteRedirect, &out.WebsiteRedirect
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObjectSpecResource.
func (in *StorageObjectSpecResource) DeepCopy() *StorageObjectSpecResource {
	if in == nil {
		return nil
	}
	out := new(StorageObjectSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObjectStatus) DeepCopyInto(out *StorageObjectStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObjectStatus.
func (in *StorageObjectStatus) DeepCopy() *StorageObjectStatus {
	if in == nil {
		return nil
	}
	out := new(StorageObjectStatus)
	in.DeepCopyInto(out)
	return out
}

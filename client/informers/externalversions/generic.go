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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "kubeform.dev/provider-linode-api/apis/domain/v1alpha1"
	firewallv1alpha1 "kubeform.dev/provider-linode-api/apis/firewall/v1alpha1"
	imagev1alpha1 "kubeform.dev/provider-linode-api/apis/image/v1alpha1"
	instancev1alpha1 "kubeform.dev/provider-linode-api/apis/instance/v1alpha1"
	lkev1alpha1 "kubeform.dev/provider-linode-api/apis/lke/v1alpha1"
	nodebalancerv1alpha1 "kubeform.dev/provider-linode-api/apis/nodebalancer/v1alpha1"
	objectv1alpha1 "kubeform.dev/provider-linode-api/apis/object/v1alpha1"
	rdnsv1alpha1 "kubeform.dev/provider-linode-api/apis/rdns/v1alpha1"
	sshkeyv1alpha1 "kubeform.dev/provider-linode-api/apis/sshkey/v1alpha1"
	stackscriptv1alpha1 "kubeform.dev/provider-linode-api/apis/stackscript/v1alpha1"
	tokenv1alpha1 "kubeform.dev/provider-linode-api/apis/token/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-linode-api/apis/user/v1alpha1"
	volumev1alpha1 "kubeform.dev/provider-linode-api/apis/volume/v1alpha1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=domain.linode.kubeform.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("domains"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Domain().V1alpha1().Domains().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("records"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Domain().V1alpha1().Records().Informer()}, nil

		// Group=firewall.linode.kubeform.com, Version=v1alpha1
	case firewallv1alpha1.SchemeGroupVersion.WithResource("firewalls"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Firewall().V1alpha1().Firewalls().Informer()}, nil

		// Group=image.linode.kubeform.com, Version=v1alpha1
	case imagev1alpha1.SchemeGroupVersion.WithResource("images"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Image().V1alpha1().Images().Informer()}, nil

		// Group=instance.linode.kubeform.com, Version=v1alpha1
	case instancev1alpha1.SchemeGroupVersion.WithResource("instances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Instance().V1alpha1().Instances().Informer()}, nil
	case instancev1alpha1.SchemeGroupVersion.WithResource("ips"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Instance().V1alpha1().Ips().Informer()}, nil

		// Group=lke.linode.kubeform.com, Version=v1alpha1
	case lkev1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Lke().V1alpha1().Clusters().Informer()}, nil

		// Group=nodebalancer.linode.kubeform.com, Version=v1alpha1
	case nodebalancerv1alpha1.SchemeGroupVersion.WithResource("configs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Nodebalancer().V1alpha1().Configs().Informer()}, nil
	case nodebalancerv1alpha1.SchemeGroupVersion.WithResource("nodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Nodebalancer().V1alpha1().Nodes().Informer()}, nil
	case nodebalancerv1alpha1.SchemeGroupVersion.WithResource("nodebalancers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Nodebalancer().V1alpha1().Nodebalancers().Informer()}, nil

		// Group=object.linode.kubeform.com, Version=v1alpha1
	case objectv1alpha1.SchemeGroupVersion.WithResource("storagebuckets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Object().V1alpha1().StorageBuckets().Informer()}, nil
	case objectv1alpha1.SchemeGroupVersion.WithResource("storagekeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Object().V1alpha1().StorageKeys().Informer()}, nil
	case objectv1alpha1.SchemeGroupVersion.WithResource("storageobjects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Object().V1alpha1().StorageObjects().Informer()}, nil

		// Group=rdns.linode.kubeform.com, Version=v1alpha1
	case rdnsv1alpha1.SchemeGroupVersion.WithResource("rdnses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rdns().V1alpha1().Rdnses().Informer()}, nil

		// Group=sshkey.linode.kubeform.com, Version=v1alpha1
	case sshkeyv1alpha1.SchemeGroupVersion.WithResource("sshkeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sshkey().V1alpha1().Sshkeys().Informer()}, nil

		// Group=stackscript.linode.kubeform.com, Version=v1alpha1
	case stackscriptv1alpha1.SchemeGroupVersion.WithResource("stackscripts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Stackscript().V1alpha1().Stackscripts().Informer()}, nil

		// Group=token.linode.kubeform.com, Version=v1alpha1
	case tokenv1alpha1.SchemeGroupVersion.WithResource("tokens"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Token().V1alpha1().Tokens().Informer()}, nil

		// Group=user.linode.kubeform.com, Version=v1alpha1
	case userv1alpha1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.User().V1alpha1().Users().Informer()}, nil

		// Group=volume.linode.kubeform.com, Version=v1alpha1
	case volumev1alpha1.SchemeGroupVersion.WithResource("volumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Volume().V1alpha1().Volumes().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
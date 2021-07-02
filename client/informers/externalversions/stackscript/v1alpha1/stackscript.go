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

package v1alpha1

import (
	"context"
	time "time"

	stackscriptv1alpha1 "kubeform.dev/provider-linode-api/apis/stackscript/v1alpha1"
	versioned "kubeform.dev/provider-linode-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-linode-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-linode-api/client/listers/stackscript/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StackscriptInformer provides access to a shared informer and lister for
// Stackscripts.
type StackscriptInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StackscriptLister
}

type stackscriptInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStackscriptInformer constructs a new informer for Stackscript type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStackscriptInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStackscriptInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStackscriptInformer constructs a new informer for Stackscript type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStackscriptInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StackscriptV1alpha1().Stackscripts(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StackscriptV1alpha1().Stackscripts(namespace).Watch(context.TODO(), options)
			},
		},
		&stackscriptv1alpha1.Stackscript{},
		resyncPeriod,
		indexers,
	)
}

func (f *stackscriptInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStackscriptInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *stackscriptInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&stackscriptv1alpha1.Stackscript{}, f.defaultInformer)
}

func (f *stackscriptInformer) Lister() v1alpha1.StackscriptLister {
	return v1alpha1.NewStackscriptLister(f.Informer().GetIndexer())
}
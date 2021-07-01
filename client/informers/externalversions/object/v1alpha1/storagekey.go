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

	objectv1alpha1 "kubeform.dev/provider-linode-api/apis/object/v1alpha1"
	versioned "kubeform.dev/provider-linode-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-linode-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-linode-api/client/listers/object/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StorageKeyInformer provides access to a shared informer and lister for
// StorageKeys.
type StorageKeyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StorageKeyLister
}

type storageKeyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStorageKeyInformer constructs a new informer for StorageKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageKeyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStorageKeyInformer constructs a new informer for StorageKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ObjectV1alpha1().StorageKeys(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ObjectV1alpha1().StorageKeys(namespace).Watch(context.TODO(), options)
			},
		},
		&objectv1alpha1.StorageKey{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageKeyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageKeyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageKeyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&objectv1alpha1.StorageKey{}, f.defaultInformer)
}

func (f *storageKeyInformer) Lister() v1alpha1.StorageKeyLister {
	return v1alpha1.NewStorageKeyLister(f.Informer().GetIndexer())
}

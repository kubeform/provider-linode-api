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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-linode-api/apis/domain/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RecordLister helps list Records.
// All objects returned here must be treated as read-only.
type RecordLister interface {
	// List lists all Records in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Record, err error)
	// Records returns an object that can list and get Records.
	Records(namespace string) RecordNamespaceLister
	RecordListerExpansion
}

// recordLister implements the RecordLister interface.
type recordLister struct {
	indexer cache.Indexer
}

// NewRecordLister returns a new RecordLister.
func NewRecordLister(indexer cache.Indexer) RecordLister {
	return &recordLister{indexer: indexer}
}

// List lists all Records in the indexer.
func (s *recordLister) List(selector labels.Selector) (ret []*v1alpha1.Record, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Record))
	})
	return ret, err
}

// Records returns an object that can list and get Records.
func (s *recordLister) Records(namespace string) RecordNamespaceLister {
	return recordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RecordNamespaceLister helps list and get Records.
// All objects returned here must be treated as read-only.
type RecordNamespaceLister interface {
	// List lists all Records in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Record, err error)
	// Get retrieves the Record from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Record, error)
	RecordNamespaceListerExpansion
}

// recordNamespaceLister implements the RecordNamespaceLister
// interface.
type recordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Records in the indexer for a given namespace.
func (s recordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Record, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Record))
	})
	return ret, err
}

// Get retrieves the Record from the indexer for a given namespace and name.
func (s recordNamespaceLister) Get(name string) (*v1alpha1.Record, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("record"), name)
	}
	return obj.(*v1alpha1.Record), nil
}

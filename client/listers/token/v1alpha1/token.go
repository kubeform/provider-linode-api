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
	v1alpha1 "kubeform.dev/provider-linode-api/apis/token/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TokenLister helps list Tokens.
// All objects returned here must be treated as read-only.
type TokenLister interface {
	// List lists all Tokens in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Token, err error)
	// Tokens returns an object that can list and get Tokens.
	Tokens(namespace string) TokenNamespaceLister
	TokenListerExpansion
}

// tokenLister implements the TokenLister interface.
type tokenLister struct {
	indexer cache.Indexer
}

// NewTokenLister returns a new TokenLister.
func NewTokenLister(indexer cache.Indexer) TokenLister {
	return &tokenLister{indexer: indexer}
}

// List lists all Tokens in the indexer.
func (s *tokenLister) List(selector labels.Selector) (ret []*v1alpha1.Token, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Token))
	})
	return ret, err
}

// Tokens returns an object that can list and get Tokens.
func (s *tokenLister) Tokens(namespace string) TokenNamespaceLister {
	return tokenNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TokenNamespaceLister helps list and get Tokens.
// All objects returned here must be treated as read-only.
type TokenNamespaceLister interface {
	// List lists all Tokens in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Token, err error)
	// Get retrieves the Token from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Token, error)
	TokenNamespaceListerExpansion
}

// tokenNamespaceLister implements the TokenNamespaceLister
// interface.
type tokenNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Tokens in the indexer for a given namespace.
func (s tokenNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Token, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Token))
	})
	return ret, err
}

// Get retrieves the Token from the indexer for a given namespace and name.
func (s tokenNamespaceLister) Get(name string) (*v1alpha1.Token, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("token"), name)
	}
	return obj.(*v1alpha1.Token), nil
}

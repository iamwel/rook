/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "github.com/rook/rook/pkg/apis/edgefs.rook.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ISGWLister helps list ISGWs.
type ISGWLister interface {
	// List lists all ISGWs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ISGW, err error)
	// ISGWs returns an object that can list and get ISGWs.
	ISGWs(namespace string) ISGWNamespaceLister
	ISGWListerExpansion
}

// iSGWLister implements the ISGWLister interface.
type iSGWLister struct {
	indexer cache.Indexer
}

// NewISGWLister returns a new ISGWLister.
func NewISGWLister(indexer cache.Indexer) ISGWLister {
	return &iSGWLister{indexer: indexer}
}

// List lists all ISGWs in the indexer.
func (s *iSGWLister) List(selector labels.Selector) (ret []*v1alpha1.ISGW, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ISGW))
	})
	return ret, err
}

// ISGWs returns an object that can list and get ISGWs.
func (s *iSGWLister) ISGWs(namespace string) ISGWNamespaceLister {
	return iSGWNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ISGWNamespaceLister helps list and get ISGWs.
type ISGWNamespaceLister interface {
	// List lists all ISGWs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ISGW, err error)
	// Get retrieves the ISGW from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ISGW, error)
	ISGWNamespaceListerExpansion
}

// iSGWNamespaceLister implements the ISGWNamespaceLister
// interface.
type iSGWNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ISGWs in the indexer for a given namespace.
func (s iSGWNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ISGW, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ISGW))
	})
	return ret, err
}

// Get retrieves the ISGW from the indexer for a given namespace and name.
func (s iSGWNamespaceLister) Get(name string) (*v1alpha1.ISGW, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("isgw"), name)
	}
	return obj.(*v1alpha1.ISGW), nil
}

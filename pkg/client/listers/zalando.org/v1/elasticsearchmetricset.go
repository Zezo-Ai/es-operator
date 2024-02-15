/*
Copyright 2024 The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/zalando-incubator/es-operator/pkg/apis/zalando.org/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ElasticsearchMetricSetLister helps list ElasticsearchMetricSets.
// All objects returned here must be treated as read-only.
type ElasticsearchMetricSetLister interface {
	// List lists all ElasticsearchMetricSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ElasticsearchMetricSet, err error)
	// ElasticsearchMetricSets returns an object that can list and get ElasticsearchMetricSets.
	ElasticsearchMetricSets(namespace string) ElasticsearchMetricSetNamespaceLister
	ElasticsearchMetricSetListerExpansion
}

// elasticsearchMetricSetLister implements the ElasticsearchMetricSetLister interface.
type elasticsearchMetricSetLister struct {
	indexer cache.Indexer
}

// NewElasticsearchMetricSetLister returns a new ElasticsearchMetricSetLister.
func NewElasticsearchMetricSetLister(indexer cache.Indexer) ElasticsearchMetricSetLister {
	return &elasticsearchMetricSetLister{indexer: indexer}
}

// List lists all ElasticsearchMetricSets in the indexer.
func (s *elasticsearchMetricSetLister) List(selector labels.Selector) (ret []*v1.ElasticsearchMetricSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ElasticsearchMetricSet))
	})
	return ret, err
}

// ElasticsearchMetricSets returns an object that can list and get ElasticsearchMetricSets.
func (s *elasticsearchMetricSetLister) ElasticsearchMetricSets(namespace string) ElasticsearchMetricSetNamespaceLister {
	return elasticsearchMetricSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ElasticsearchMetricSetNamespaceLister helps list and get ElasticsearchMetricSets.
// All objects returned here must be treated as read-only.
type ElasticsearchMetricSetNamespaceLister interface {
	// List lists all ElasticsearchMetricSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ElasticsearchMetricSet, err error)
	// Get retrieves the ElasticsearchMetricSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ElasticsearchMetricSet, error)
	ElasticsearchMetricSetNamespaceListerExpansion
}

// elasticsearchMetricSetNamespaceLister implements the ElasticsearchMetricSetNamespaceLister
// interface.
type elasticsearchMetricSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ElasticsearchMetricSets in the indexer for a given namespace.
func (s elasticsearchMetricSetNamespaceLister) List(selector labels.Selector) (ret []*v1.ElasticsearchMetricSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ElasticsearchMetricSet))
	})
	return ret, err
}

// Get retrieves the ElasticsearchMetricSet from the indexer for a given namespace and name.
func (s elasticsearchMetricSetNamespaceLister) Get(name string) (*v1.ElasticsearchMetricSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("elasticsearchmetricset"), name)
	}
	return obj.(*v1.ElasticsearchMetricSet), nil
}

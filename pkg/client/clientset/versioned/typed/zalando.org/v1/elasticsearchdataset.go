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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/zalando-incubator/es-operator/pkg/apis/zalando.org/v1"
	scheme "github.com/zalando-incubator/es-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ElasticsearchDataSetsGetter has a method to return a ElasticsearchDataSetInterface.
// A group's client should implement this interface.
type ElasticsearchDataSetsGetter interface {
	ElasticsearchDataSets(namespace string) ElasticsearchDataSetInterface
}

// ElasticsearchDataSetInterface has methods to work with ElasticsearchDataSet resources.
type ElasticsearchDataSetInterface interface {
	Create(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.CreateOptions) (*v1.ElasticsearchDataSet, error)
	Update(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.UpdateOptions) (*v1.ElasticsearchDataSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, elasticsearchDataSet *v1.ElasticsearchDataSet, opts metav1.UpdateOptions) (*v1.ElasticsearchDataSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ElasticsearchDataSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ElasticsearchDataSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ElasticsearchDataSet, err error)
	ElasticsearchDataSetExpansion
}

// elasticsearchDataSets implements ElasticsearchDataSetInterface
type elasticsearchDataSets struct {
	*gentype.ClientWithList[*v1.ElasticsearchDataSet, *v1.ElasticsearchDataSetList]
}

// newElasticsearchDataSets returns a ElasticsearchDataSets
func newElasticsearchDataSets(c *ZalandoV1Client, namespace string) *elasticsearchDataSets {
	return &elasticsearchDataSets{
		gentype.NewClientWithList[*v1.ElasticsearchDataSet, *v1.ElasticsearchDataSetList](
			"elasticsearchdatasets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.ElasticsearchDataSet { return &v1.ElasticsearchDataSet{} },
			func() *v1.ElasticsearchDataSetList { return &v1.ElasticsearchDataSetList{} }),
	}
}

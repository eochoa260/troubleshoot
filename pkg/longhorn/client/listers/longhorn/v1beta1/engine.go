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

package v1beta1

import (
	v1beta1 "github.com/replicatedhq/troubleshoot/pkg/longhorn/apis/longhorn/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EngineLister helps list Engines.
type EngineLister interface {
	// List lists all Engines in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.Engine, err error)
	// Engines returns an object that can list and get Engines.
	Engines(namespace string) EngineNamespaceLister
	EngineListerExpansion
}

// engineLister implements the EngineLister interface.
type engineLister struct {
	indexer cache.Indexer
}

// NewEngineLister returns a new EngineLister.
func NewEngineLister(indexer cache.Indexer) EngineLister {
	return &engineLister{indexer: indexer}
}

// List lists all Engines in the indexer.
func (s *engineLister) List(selector labels.Selector) (ret []*v1beta1.Engine, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Engine))
	})
	return ret, err
}

// Engines returns an object that can list and get Engines.
func (s *engineLister) Engines(namespace string) EngineNamespaceLister {
	return engineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EngineNamespaceLister helps list and get Engines.
type EngineNamespaceLister interface {
	// List lists all Engines in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.Engine, err error)
	// Get retrieves the Engine from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.Engine, error)
	EngineNamespaceListerExpansion
}

// engineNamespaceLister implements the EngineNamespaceLister
// interface.
type engineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Engines in the indexer for a given namespace.
func (s engineNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Engine, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Engine))
	})
	return ret, err
}

// Get retrieves the Engine from the indexer for a given namespace and name.
func (s engineNamespaceLister) Get(name string) (*v1beta1.Engine, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("engine"), name)
	}
	return obj.(*v1beta1.Engine), nil
}

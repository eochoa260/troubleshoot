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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	longhornv1beta1 "github.com/replicatedhq/troubleshoot/pkg/longhorn/apis/longhorn/v1beta1"
	versioned "github.com/replicatedhq/troubleshoot/pkg/longhorn/client/clientset/versioned"
	internalinterfaces "github.com/replicatedhq/troubleshoot/pkg/longhorn/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/replicatedhq/troubleshoot/pkg/longhorn/client/listers/longhorn/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReplicaInformer provides access to a shared informer and lister for
// Replicas.
type ReplicaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ReplicaLister
}

type replicaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReplicaInformer constructs a new informer for Replica type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReplicaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReplicaInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReplicaInformer constructs a new informer for Replica type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReplicaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LonghornV1beta1().Replicas(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LonghornV1beta1().Replicas(namespace).Watch(context.TODO(), options)
			},
		},
		&longhornv1beta1.Replica{},
		resyncPeriod,
		indexers,
	)
}

func (f *replicaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReplicaInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *replicaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&longhornv1beta1.Replica{}, f.defaultInformer)
}

func (f *replicaInformer) Lister() v1beta1.ReplicaLister {
	return v1beta1.NewReplicaLister(f.Informer().GetIndexer())
}

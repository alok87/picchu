// Copyright © 2019 A Medium Corporation.
// Licensed under the Apache License, Version 2.0; see the NOTICE file.

// Code generated by informer. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	picchuv1alpha1 "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1"
	client "go.medium.engineering/picchu/pkg/client"
	internalinterfaces "go.medium.engineering/picchu/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "go.medium.engineering/picchu/pkg/client/listers/picchu/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterInformer provides access to a shared informer and lister for
// Clusters.
type ClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterLister
}

type clusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterInformer constructs a new informer for Cluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterInformer(client client.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterInformer constructs a new informer for Cluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterInformer(client client.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PicchuV1alpha1().Clusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PicchuV1alpha1().Clusters(namespace).Watch(options)
			},
		},
		&picchuv1alpha1.Cluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterInformer) defaultInformer(client client.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&picchuv1alpha1.Cluster{}, f.defaultInformer)
}

func (f *clusterInformer) Lister() v1alpha1.ClusterLister {
	return v1alpha1.NewClusterLister(f.Informer().GetIndexer())
}
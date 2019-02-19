// Copyright © 2019 A Medium Corporation.
// Licensed under the Apache License, Version 2.0; see the LICENSE file.

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

// RevisionInformer provides access to a shared informer and lister for
// Revisions.
type RevisionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RevisionLister
}

type revisionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRevisionInformer constructs a new informer for Revision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRevisionInformer(client client.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRevisionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRevisionInformer constructs a new informer for Revision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRevisionInformer(client client.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PicchuV1alpha1().Revisions(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PicchuV1alpha1().Revisions(namespace).Watch(options)
			},
		},
		&picchuv1alpha1.Revision{},
		resyncPeriod,
		indexers,
	)
}

func (f *revisionInformer) defaultInformer(client client.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRevisionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *revisionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&picchuv1alpha1.Revision{}, f.defaultInformer)
}

func (f *revisionInformer) Lister() v1alpha1.RevisionLister {
	return v1alpha1.NewRevisionLister(f.Informer().GetIndexer())
}

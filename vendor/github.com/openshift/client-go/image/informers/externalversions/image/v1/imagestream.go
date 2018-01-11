// This file was automatically generated by informer-gen

package v1

import (
	image_v1 "github.com/openshift/api/image/v1"
	versioned "github.com/openshift/client-go/image/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/image/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/image/listers/image/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ImageStreamInformer provides access to a shared informer and lister for
// ImageStreams.
type ImageStreamInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ImageStreamLister
}

type imageStreamInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewImageStreamInformer constructs a new informer for ImageStream type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewImageStreamInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.ImageV1().ImageStreams(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.ImageV1().ImageStreams(namespace).Watch(options)
			},
		},
		&image_v1.ImageStream{},
		resyncPeriod,
		indexers,
	)
}

func defaultImageStreamInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewImageStreamInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *imageStreamInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&image_v1.ImageStream{}, defaultImageStreamInformer)
}

func (f *imageStreamInformer) Lister() v1.ImageStreamLister {
	return v1.NewImageStreamLister(f.Informer().GetIndexer())
}

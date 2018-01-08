// This file was automatically generated by informer-gen

package v1

import (
	authorization_v1 "github.com/openshift/api/authorization/v1"
	versioned "github.com/openshift/client-go/authorization/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/authorization/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/authorization/listers/authorization/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// RoleBindingRestrictionInformer provides access to a shared informer and lister for
// RoleBindingRestrictions.
type RoleBindingRestrictionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RoleBindingRestrictionLister
}

type roleBindingRestrictionInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewRoleBindingRestrictionInformer constructs a new informer for RoleBindingRestriction type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRoleBindingRestrictionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AuthorizationV1().RoleBindingRestrictions(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AuthorizationV1().RoleBindingRestrictions(namespace).Watch(options)
			},
		},
		&authorization_v1.RoleBindingRestriction{},
		resyncPeriod,
		indexers,
	)
}

func defaultRoleBindingRestrictionInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewRoleBindingRestrictionInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *roleBindingRestrictionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authorization_v1.RoleBindingRestriction{}, defaultRoleBindingRestrictionInformer)
}

func (f *roleBindingRestrictionInformer) Lister() v1.RoleBindingRestrictionLister {
	return v1.NewRoleBindingRestrictionLister(f.Informer().GetIndexer())
}

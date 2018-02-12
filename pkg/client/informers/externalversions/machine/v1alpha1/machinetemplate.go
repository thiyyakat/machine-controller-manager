// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	machine_v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	versioned "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned"
	internalinterfaces "github.com/gardener/machine-controller-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/client/listers/machine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MachineTemplateInformer provides access to a shared informer and lister for
// MachineTemplates.
type MachineTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineTemplateLister
}

type machineTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMachineTemplateInformer constructs a new informer for MachineTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMachineTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMachineTemplateInformer constructs a new informer for MachineTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMachineTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1alpha1().MachineTemplates(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1alpha1().MachineTemplates(namespace).Watch(options)
			},
		},
		&machine_v1alpha1.MachineTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *machineTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMachineTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *machineTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machine_v1alpha1.MachineTemplate{}, f.defaultInformer)
}

func (f *machineTemplateInformer) Lister() v1alpha1.MachineTemplateLister {
	return v1alpha1.NewMachineTemplateLister(f.Informer().GetIndexer())
}

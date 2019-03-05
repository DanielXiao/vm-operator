/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	vmoperatorv1alpha1 "vmware.com/kubevsphere/pkg/apis/vmoperator/v1alpha1"
	clientset "vmware.com/kubevsphere/pkg/client/clientset_generated/clientset"
	internalinterfaces "vmware.com/kubevsphere/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "vmware.com/kubevsphere/pkg/client/listers_generated/vmoperator/v1alpha1"
)

// VirtualMachineClassInformer provides access to a shared informer and lister for
// VirtualMachineClasses.
type VirtualMachineClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VirtualMachineClassLister
}

type virtualMachineClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualMachineClassInformer constructs a new informer for VirtualMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualMachineClassInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualMachineClassInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualMachineClassInformer constructs a new informer for VirtualMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualMachineClassInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VmoperatorV1alpha1().VirtualMachineClasses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VmoperatorV1alpha1().VirtualMachineClasses(namespace).Watch(options)
			},
		},
		&vmoperatorv1alpha1.VirtualMachineClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualMachineClassInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualMachineClassInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualMachineClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&vmoperatorv1alpha1.VirtualMachineClass{}, f.defaultInformer)
}

func (f *virtualMachineClassInformer) Lister() v1alpha1.VirtualMachineClassLister {
	return v1alpha1.NewVirtualMachineClassLister(f.Informer().GetIndexer())
}

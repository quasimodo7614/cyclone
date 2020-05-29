/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package informers

import (
	v1alpha1 "github.com/caicloud/cyclone/pkg/apis/cyclone/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers"
	cache "k8s.io/client-go/tools/cache"
)

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error) {
	switch resource {
	// Group=cyclone.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("executionclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cyclone().V1alpha1().ExecutionClusters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cyclone().V1alpha1().Projects().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("resources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cyclone().V1alpha1().Resources().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("stages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cyclone().V1alpha1().Stages().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("workflows"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cyclone().V1alpha1().Workflows().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("workflowruns"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cyclone().V1alpha1().WorkflowRuns().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("workflowtriggers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cyclone().V1alpha1().WorkflowTriggers().Informer()}, nil

	}

	return f.SharedInformerFactory.ForResource(resource)
}

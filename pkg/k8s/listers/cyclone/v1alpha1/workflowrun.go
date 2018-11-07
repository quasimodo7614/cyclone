/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/cyclone/pkg/apis/cyclone/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkflowRunLister helps list WorkflowRuns.
type WorkflowRunLister interface {
	// List lists all WorkflowRuns in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WorkflowRun, err error)
	// WorkflowRuns returns an object that can list and get WorkflowRuns.
	WorkflowRuns(namespace string) WorkflowRunNamespaceLister
	WorkflowRunListerExpansion
}

// workflowRunLister implements the WorkflowRunLister interface.
type workflowRunLister struct {
	indexer cache.Indexer
}

// NewWorkflowRunLister returns a new WorkflowRunLister.
func NewWorkflowRunLister(indexer cache.Indexer) WorkflowRunLister {
	return &workflowRunLister{indexer: indexer}
}

// List lists all WorkflowRuns in the indexer.
func (s *workflowRunLister) List(selector labels.Selector) (ret []*v1alpha1.WorkflowRun, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkflowRun))
	})
	return ret, err
}

// WorkflowRuns returns an object that can list and get WorkflowRuns.
func (s *workflowRunLister) WorkflowRuns(namespace string) WorkflowRunNamespaceLister {
	return workflowRunNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkflowRunNamespaceLister helps list and get WorkflowRuns.
type WorkflowRunNamespaceLister interface {
	// List lists all WorkflowRuns in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WorkflowRun, err error)
	// Get retrieves the WorkflowRun from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WorkflowRun, error)
	WorkflowRunNamespaceListerExpansion
}

// workflowRunNamespaceLister implements the WorkflowRunNamespaceLister
// interface.
type workflowRunNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorkflowRuns in the indexer for a given namespace.
func (s workflowRunNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WorkflowRun, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkflowRun))
	})
	return ret, err
}

// Get retrieves the WorkflowRun from the indexer for a given namespace and name.
func (s workflowRunNamespaceLister) Get(name string) (*v1alpha1.WorkflowRun, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workflowrun"), name)
	}
	return obj.(*v1alpha1.WorkflowRun), nil
}

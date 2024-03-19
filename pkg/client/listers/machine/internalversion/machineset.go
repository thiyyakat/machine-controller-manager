// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineSetLister helps list MachineSets.
// All objects returned here must be treated as read-only.
type MachineSetLister interface {
	// List lists all MachineSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*machine.MachineSet, err error)
	// MachineSets returns an object that can list and get MachineSets.
	MachineSets(namespace string) MachineSetNamespaceLister
	MachineSetListerExpansion
}

// machineSetLister implements the MachineSetLister interface.
type machineSetLister struct {
	indexer cache.Indexer
}

// NewMachineSetLister returns a new MachineSetLister.
func NewMachineSetLister(indexer cache.Indexer) MachineSetLister {
	return &machineSetLister{indexer: indexer}
}

// List lists all MachineSets in the indexer.
func (s *machineSetLister) List(selector labels.Selector) (ret []*machine.MachineSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.MachineSet))
	})
	return ret, err
}

// MachineSets returns an object that can list and get MachineSets.
func (s *machineSetLister) MachineSets(namespace string) MachineSetNamespaceLister {
	return machineSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineSetNamespaceLister helps list and get MachineSets.
// All objects returned here must be treated as read-only.
type MachineSetNamespaceLister interface {
	// List lists all MachineSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*machine.MachineSet, err error)
	// Get retrieves the MachineSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*machine.MachineSet, error)
	MachineSetNamespaceListerExpansion
}

// machineSetNamespaceLister implements the MachineSetNamespaceLister
// interface.
type machineSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineSets in the indexer for a given namespace.
func (s machineSetNamespaceLister) List(selector labels.Selector) (ret []*machine.MachineSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.MachineSet))
	})
	return ret, err
}

// Get retrieves the MachineSet from the indexer for a given namespace and name.
func (s machineSetNamespaceLister) Get(name string) (*machine.MachineSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("machineset"), name)
	}
	return obj.(*machine.MachineSet), nil
}

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineTemplates implements MachineTemplateInterface
type FakeMachineTemplates struct {
	Fake *FakeMachine
	ns   string
}

var machinetemplatesResource = schema.GroupVersionResource{Group: "machine.sapcloud.io", Version: "", Resource: "machinetemplates"}

var machinetemplatesKind = schema.GroupVersionKind{Group: "machine.sapcloud.io", Version: "", Kind: "MachineTemplate"}

// Get takes name of the machineTemplate, and returns the corresponding machineTemplate object, and an error if there is any.
func (c *FakeMachineTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *machine.MachineTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machinetemplatesResource, c.ns, name), &machine.MachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.MachineTemplate), err
}

// List takes label and field selectors, and returns the list of MachineTemplates that match those selectors.
func (c *FakeMachineTemplates) List(ctx context.Context, opts v1.ListOptions) (result *machine.MachineTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machinetemplatesResource, machinetemplatesKind, c.ns, opts), &machine.MachineTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machine.MachineTemplateList{ListMeta: obj.(*machine.MachineTemplateList).ListMeta}
	for _, item := range obj.(*machine.MachineTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineTemplates.
func (c *FakeMachineTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machinetemplatesResource, c.ns, opts))

}

// Create takes the representation of a machineTemplate and creates it.  Returns the server's representation of the machineTemplate, and an error, if there is any.
func (c *FakeMachineTemplates) Create(ctx context.Context, machineTemplate *machine.MachineTemplate, opts v1.CreateOptions) (result *machine.MachineTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machinetemplatesResource, c.ns, machineTemplate), &machine.MachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.MachineTemplate), err
}

// Update takes the representation of a machineTemplate and updates it. Returns the server's representation of the machineTemplate, and an error, if there is any.
func (c *FakeMachineTemplates) Update(ctx context.Context, machineTemplate *machine.MachineTemplate, opts v1.UpdateOptions) (result *machine.MachineTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machinetemplatesResource, c.ns, machineTemplate), &machine.MachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.MachineTemplate), err
}

// Delete takes name of the machineTemplate and deletes it. Returns an error if one occurs.
func (c *FakeMachineTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(machinetemplatesResource, c.ns, name, opts), &machine.MachineTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machinetemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &machine.MachineTemplateList{})
	return err
}

// Patch applies the patch and returns the patched machineTemplate.
func (c *FakeMachineTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *machine.MachineTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinetemplatesResource, c.ns, name, pt, data, subresources...), &machine.MachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.MachineTemplate), err
}

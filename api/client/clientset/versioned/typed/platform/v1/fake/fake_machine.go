/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakeMachines implements MachineInterface
type FakeMachines struct {
	Fake *FakePlatformV1
}

var machinesResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "machines"}

var machinesKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "Machine"}

// Get takes name of the machine, and returns the corresponding machine object, and an error if there is any.
func (c *FakeMachines) Get(name string, options v1.GetOptions) (result *platformv1.Machine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(machinesResource, name), &platformv1.Machine{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Machine), err
}

// List takes label and field selectors, and returns the list of Machines that match those selectors.
func (c *FakeMachines) List(opts v1.ListOptions) (result *platformv1.MachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(machinesResource, machinesKind, opts), &platformv1.MachineList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.MachineList{ListMeta: obj.(*platformv1.MachineList).ListMeta}
	for _, item := range obj.(*platformv1.MachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machines.
func (c *FakeMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(machinesResource, opts))
}

// Create takes the representation of a machine and creates it.  Returns the server's representation of the machine, and an error, if there is any.
func (c *FakeMachines) Create(machine *platformv1.Machine) (result *platformv1.Machine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(machinesResource, machine), &platformv1.Machine{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Machine), err
}

// Update takes the representation of a machine and updates it. Returns the server's representation of the machine, and an error, if there is any.
func (c *FakeMachines) Update(machine *platformv1.Machine) (result *platformv1.Machine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(machinesResource, machine), &platformv1.Machine{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Machine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachines) UpdateStatus(machine *platformv1.Machine) (*platformv1.Machine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(machinesResource, "status", machine), &platformv1.Machine{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Machine), err
}

// Delete takes name of the machine and deletes it. Returns an error if one occurs.
func (c *FakeMachines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(machinesResource, name), &platformv1.Machine{})
	return err
}

// Patch applies the patch and returns the patched machine.
func (c *FakeMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platformv1.Machine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(machinesResource, name, pt, data, subresources...), &platformv1.Machine{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Machine), err
}

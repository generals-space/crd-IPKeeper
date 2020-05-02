/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	ipkeeperv1 "github.com/generals-space/crd-ipkeeper/pkg/apis/ipkeeper/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStaticIPs implements StaticIPInterface
type FakeStaticIPs struct {
	Fake *FakeIpkeeperV1
	ns   string
}

var staticipsResource = schema.GroupVersionResource{Group: "ipkeeper.generals.space", Version: "v1", Resource: "staticips"}

var staticipsKind = schema.GroupVersionKind{Group: "ipkeeper.generals.space", Version: "v1", Kind: "StaticIP"}

// Get takes name of the staticIP, and returns the corresponding staticIP object, and an error if there is any.
func (c *FakeStaticIPs) Get(name string, options v1.GetOptions) (result *ipkeeperv1.StaticIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(staticipsResource, c.ns, name), &ipkeeperv1.StaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ipkeeperv1.StaticIP), err
}

// List takes label and field selectors, and returns the list of StaticIPs that match those selectors.
func (c *FakeStaticIPs) List(opts v1.ListOptions) (result *ipkeeperv1.StaticIPList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(staticipsResource, staticipsKind, c.ns, opts), &ipkeeperv1.StaticIPList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &ipkeeperv1.StaticIPList{ListMeta: obj.(*ipkeeperv1.StaticIPList).ListMeta}
	for _, item := range obj.(*ipkeeperv1.StaticIPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested staticIPs.
func (c *FakeStaticIPs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(staticipsResource, c.ns, opts))

}

// Create takes the representation of a staticIP and creates it.  Returns the server's representation of the staticIP, and an error, if there is any.
func (c *FakeStaticIPs) Create(staticIP *ipkeeperv1.StaticIP) (result *ipkeeperv1.StaticIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(staticipsResource, c.ns, staticIP), &ipkeeperv1.StaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ipkeeperv1.StaticIP), err
}

// Update takes the representation of a staticIP and updates it. Returns the server's representation of the staticIP, and an error, if there is any.
func (c *FakeStaticIPs) Update(staticIP *ipkeeperv1.StaticIP) (result *ipkeeperv1.StaticIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(staticipsResource, c.ns, staticIP), &ipkeeperv1.StaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ipkeeperv1.StaticIP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStaticIPs) UpdateStatus(staticIP *ipkeeperv1.StaticIP) (*ipkeeperv1.StaticIP, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(staticipsResource, "status", c.ns, staticIP), &ipkeeperv1.StaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ipkeeperv1.StaticIP), err
}

// Delete takes name of the staticIP and deletes it. Returns an error if one occurs.
func (c *FakeStaticIPs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(staticipsResource, c.ns, name), &ipkeeperv1.StaticIP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStaticIPs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(staticipsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &ipkeeperv1.StaticIPList{})
	return err
}

// Patch applies the patch and returns the patched staticIP.
func (c *FakeStaticIPs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *ipkeeperv1.StaticIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(staticipsResource, c.ns, name, pt, data, subresources...), &ipkeeperv1.StaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*ipkeeperv1.StaticIP), err
}
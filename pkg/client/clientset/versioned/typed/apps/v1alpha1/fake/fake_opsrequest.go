/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

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
	"context"

	v1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOpsRequests implements OpsRequestInterface
type FakeOpsRequests struct {
	Fake *FakeAppsV1alpha1
	ns   string
}

var opsrequestsResource = v1alpha1.SchemeGroupVersion.WithResource("opsrequests")

var opsrequestsKind = v1alpha1.SchemeGroupVersion.WithKind("OpsRequest")

// Get takes name of the opsRequest, and returns the corresponding opsRequest object, and an error if there is any.
func (c *FakeOpsRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OpsRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(opsrequestsResource, c.ns, name), &v1alpha1.OpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsRequest), err
}

// List takes label and field selectors, and returns the list of OpsRequests that match those selectors.
func (c *FakeOpsRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OpsRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(opsrequestsResource, opsrequestsKind, c.ns, opts), &v1alpha1.OpsRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsRequestList{ListMeta: obj.(*v1alpha1.OpsRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsRequests.
func (c *FakeOpsRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(opsrequestsResource, c.ns, opts))

}

// Create takes the representation of a opsRequest and creates it.  Returns the server's representation of the opsRequest, and an error, if there is any.
func (c *FakeOpsRequests) Create(ctx context.Context, opsRequest *v1alpha1.OpsRequest, opts v1.CreateOptions) (result *v1alpha1.OpsRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(opsrequestsResource, c.ns, opsRequest), &v1alpha1.OpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsRequest), err
}

// Update takes the representation of a opsRequest and updates it. Returns the server's representation of the opsRequest, and an error, if there is any.
func (c *FakeOpsRequests) Update(ctx context.Context, opsRequest *v1alpha1.OpsRequest, opts v1.UpdateOptions) (result *v1alpha1.OpsRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(opsrequestsResource, c.ns, opsRequest), &v1alpha1.OpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsRequests) UpdateStatus(ctx context.Context, opsRequest *v1alpha1.OpsRequest, opts v1.UpdateOptions) (*v1alpha1.OpsRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(opsrequestsResource, "status", c.ns, opsRequest), &v1alpha1.OpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsRequest), err
}

// Delete takes name of the opsRequest and deletes it. Returns an error if one occurs.
func (c *FakeOpsRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(opsrequestsResource, c.ns, name, opts), &v1alpha1.OpsRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(opsrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsRequestList{})
	return err
}

// Patch applies the patch and returns the patched opsRequest.
func (c *FakeOpsRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OpsRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(opsrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsRequest), err
}

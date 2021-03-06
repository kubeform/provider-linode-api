/*
Copyright AppsCode Inc. and Contributors

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

	v1alpha1 "kubeform.dev/provider-linode-api/apis/stackscript/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStackscripts implements StackscriptInterface
type FakeStackscripts struct {
	Fake *FakeStackscriptV1alpha1
	ns   string
}

var stackscriptsResource = schema.GroupVersionResource{Group: "stackscript.linode.kubeform.com", Version: "v1alpha1", Resource: "stackscripts"}

var stackscriptsKind = schema.GroupVersionKind{Group: "stackscript.linode.kubeform.com", Version: "v1alpha1", Kind: "Stackscript"}

// Get takes name of the stackscript, and returns the corresponding stackscript object, and an error if there is any.
func (c *FakeStackscripts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Stackscript, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(stackscriptsResource, c.ns, name), &v1alpha1.Stackscript{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stackscript), err
}

// List takes label and field selectors, and returns the list of Stackscripts that match those selectors.
func (c *FakeStackscripts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StackscriptList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(stackscriptsResource, stackscriptsKind, c.ns, opts), &v1alpha1.StackscriptList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StackscriptList{ListMeta: obj.(*v1alpha1.StackscriptList).ListMeta}
	for _, item := range obj.(*v1alpha1.StackscriptList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested stackscripts.
func (c *FakeStackscripts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(stackscriptsResource, c.ns, opts))

}

// Create takes the representation of a stackscript and creates it.  Returns the server's representation of the stackscript, and an error, if there is any.
func (c *FakeStackscripts) Create(ctx context.Context, stackscript *v1alpha1.Stackscript, opts v1.CreateOptions) (result *v1alpha1.Stackscript, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(stackscriptsResource, c.ns, stackscript), &v1alpha1.Stackscript{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stackscript), err
}

// Update takes the representation of a stackscript and updates it. Returns the server's representation of the stackscript, and an error, if there is any.
func (c *FakeStackscripts) Update(ctx context.Context, stackscript *v1alpha1.Stackscript, opts v1.UpdateOptions) (result *v1alpha1.Stackscript, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(stackscriptsResource, c.ns, stackscript), &v1alpha1.Stackscript{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stackscript), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStackscripts) UpdateStatus(ctx context.Context, stackscript *v1alpha1.Stackscript, opts v1.UpdateOptions) (*v1alpha1.Stackscript, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(stackscriptsResource, "status", c.ns, stackscript), &v1alpha1.Stackscript{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stackscript), err
}

// Delete takes name of the stackscript and deletes it. Returns an error if one occurs.
func (c *FakeStackscripts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(stackscriptsResource, c.ns, name), &v1alpha1.Stackscript{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStackscripts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(stackscriptsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StackscriptList{})
	return err
}

// Patch applies the patch and returns the patched stackscript.
func (c *FakeStackscripts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Stackscript, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(stackscriptsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Stackscript{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stackscript), err
}

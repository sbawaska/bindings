/*
Copyright 2019 the original author or authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/projectriff/bindings/pkg/apis/bindings/v1alpha1"
	scheme "github.com/projectriff/bindings/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FrogBindingsGetter has a method to return a FrogBindingInterface.
// A group's client should implement this interface.
type FrogBindingsGetter interface {
	FrogBindings(namespace string) FrogBindingInterface
}

// FrogBindingInterface has methods to work with FrogBinding resources.
type FrogBindingInterface interface {
	Create(*v1alpha1.FrogBinding) (*v1alpha1.FrogBinding, error)
	Update(*v1alpha1.FrogBinding) (*v1alpha1.FrogBinding, error)
	UpdateStatus(*v1alpha1.FrogBinding) (*v1alpha1.FrogBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FrogBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.FrogBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FrogBinding, err error)
	FrogBindingExpansion
}

// frogBindings implements FrogBindingInterface
type frogBindings struct {
	client rest.Interface
	ns     string
}

// newFrogBindings returns a FrogBindings
func newFrogBindings(c *BindingsV1alpha1Client, namespace string) *frogBindings {
	return &frogBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the frogBinding, and returns the corresponding frogBinding object, and an error if there is any.
func (c *frogBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.FrogBinding, err error) {
	result = &v1alpha1.FrogBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("frogbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FrogBindings that match those selectors.
func (c *frogBindings) List(opts v1.ListOptions) (result *v1alpha1.FrogBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FrogBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("frogbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested frogBindings.
func (c *frogBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("frogbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a frogBinding and creates it.  Returns the server's representation of the frogBinding, and an error, if there is any.
func (c *frogBindings) Create(frogBinding *v1alpha1.FrogBinding) (result *v1alpha1.FrogBinding, err error) {
	result = &v1alpha1.FrogBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("frogbindings").
		Body(frogBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a frogBinding and updates it. Returns the server's representation of the frogBinding, and an error, if there is any.
func (c *frogBindings) Update(frogBinding *v1alpha1.FrogBinding) (result *v1alpha1.FrogBinding, err error) {
	result = &v1alpha1.FrogBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("frogbindings").
		Name(frogBinding.Name).
		Body(frogBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *frogBindings) UpdateStatus(frogBinding *v1alpha1.FrogBinding) (result *v1alpha1.FrogBinding, err error) {
	result = &v1alpha1.FrogBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("frogbindings").
		Name(frogBinding.Name).
		SubResource("status").
		Body(frogBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the frogBinding and deletes it. Returns an error if one occurs.
func (c *frogBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("frogbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *frogBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("frogbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched frogBinding.
func (c *frogBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FrogBinding, err error) {
	result = &v1alpha1.FrogBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("frogbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

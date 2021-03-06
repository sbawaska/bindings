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

// ProvisionedServicesGetter has a method to return a ProvisionedServiceInterface.
// A group's client should implement this interface.
type ProvisionedServicesGetter interface {
	ProvisionedServices(namespace string) ProvisionedServiceInterface
}

// ProvisionedServiceInterface has methods to work with ProvisionedService resources.
type ProvisionedServiceInterface interface {
	Create(*v1alpha1.ProvisionedService) (*v1alpha1.ProvisionedService, error)
	Update(*v1alpha1.ProvisionedService) (*v1alpha1.ProvisionedService, error)
	UpdateStatus(*v1alpha1.ProvisionedService) (*v1alpha1.ProvisionedService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ProvisionedService, error)
	List(opts v1.ListOptions) (*v1alpha1.ProvisionedServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProvisionedService, err error)
	ProvisionedServiceExpansion
}

// provisionedServices implements ProvisionedServiceInterface
type provisionedServices struct {
	client rest.Interface
	ns     string
}

// newProvisionedServices returns a ProvisionedServices
func newProvisionedServices(c *BindingsV1alpha1Client, namespace string) *provisionedServices {
	return &provisionedServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the provisionedService, and returns the corresponding provisionedService object, and an error if there is any.
func (c *provisionedServices) Get(name string, options v1.GetOptions) (result *v1alpha1.ProvisionedService, err error) {
	result = &v1alpha1.ProvisionedService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("provisionedservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProvisionedServices that match those selectors.
func (c *provisionedServices) List(opts v1.ListOptions) (result *v1alpha1.ProvisionedServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ProvisionedServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("provisionedservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested provisionedServices.
func (c *provisionedServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("provisionedservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a provisionedService and creates it.  Returns the server's representation of the provisionedService, and an error, if there is any.
func (c *provisionedServices) Create(provisionedService *v1alpha1.ProvisionedService) (result *v1alpha1.ProvisionedService, err error) {
	result = &v1alpha1.ProvisionedService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("provisionedservices").
		Body(provisionedService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a provisionedService and updates it. Returns the server's representation of the provisionedService, and an error, if there is any.
func (c *provisionedServices) Update(provisionedService *v1alpha1.ProvisionedService) (result *v1alpha1.ProvisionedService, err error) {
	result = &v1alpha1.ProvisionedService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("provisionedservices").
		Name(provisionedService.Name).
		Body(provisionedService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *provisionedServices) UpdateStatus(provisionedService *v1alpha1.ProvisionedService) (result *v1alpha1.ProvisionedService, err error) {
	result = &v1alpha1.ProvisionedService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("provisionedservices").
		Name(provisionedService.Name).
		SubResource("status").
		Body(provisionedService).
		Do().
		Into(result)
	return
}

// Delete takes name of the provisionedService and deletes it. Returns an error if one occurs.
func (c *provisionedServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("provisionedservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *provisionedServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("provisionedservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched provisionedService.
func (c *provisionedServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProvisionedService, err error) {
	result = &v1alpha1.ProvisionedService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("provisionedservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

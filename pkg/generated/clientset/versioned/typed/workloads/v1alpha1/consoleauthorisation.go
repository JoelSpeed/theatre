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

package v1alpha1

import (
	"context"
	"time"

	scheme "github.com/gocardless/theatre/pkg/generated/clientset/versioned/scheme"
	v1alpha1 "github.com/gocardless/theatre/v3/apis/workloads/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConsoleAuthorisationsGetter has a method to return a ConsoleAuthorisationInterface.
// A group's client should implement this interface.
type ConsoleAuthorisationsGetter interface {
	ConsoleAuthorisations(namespace string) ConsoleAuthorisationInterface
}

// ConsoleAuthorisationInterface has methods to work with ConsoleAuthorisation resources.
type ConsoleAuthorisationInterface interface {
	Create(ctx context.Context, consoleAuthorisation *v1alpha1.ConsoleAuthorisation, opts v1.CreateOptions) (*v1alpha1.ConsoleAuthorisation, error)
	Update(ctx context.Context, consoleAuthorisation *v1alpha1.ConsoleAuthorisation, opts v1.UpdateOptions) (*v1alpha1.ConsoleAuthorisation, error)
	UpdateStatus(ctx context.Context, consoleAuthorisation *v1alpha1.ConsoleAuthorisation, opts v1.UpdateOptions) (*v1alpha1.ConsoleAuthorisation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ConsoleAuthorisation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ConsoleAuthorisationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConsoleAuthorisation, err error)
	ConsoleAuthorisationExpansion
}

// consoleAuthorisations implements ConsoleAuthorisationInterface
type consoleAuthorisations struct {
	client rest.Interface
	ns     string
}

// newConsoleAuthorisations returns a ConsoleAuthorisations
func newConsoleAuthorisations(c *WorkloadsV1alpha1Client, namespace string) *consoleAuthorisations {
	return &consoleAuthorisations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the consoleAuthorisation, and returns the corresponding consoleAuthorisation object, and an error if there is any.
func (c *consoleAuthorisations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConsoleAuthorisation, err error) {
	result = &v1alpha1.ConsoleAuthorisation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("consoleauthorisations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConsoleAuthorisations that match those selectors.
func (c *consoleAuthorisations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConsoleAuthorisationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ConsoleAuthorisationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("consoleauthorisations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested consoleAuthorisations.
func (c *consoleAuthorisations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("consoleauthorisations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a consoleAuthorisation and creates it.  Returns the server's representation of the consoleAuthorisation, and an error, if there is any.
func (c *consoleAuthorisations) Create(ctx context.Context, consoleAuthorisation *v1alpha1.ConsoleAuthorisation, opts v1.CreateOptions) (result *v1alpha1.ConsoleAuthorisation, err error) {
	result = &v1alpha1.ConsoleAuthorisation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("consoleauthorisations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(consoleAuthorisation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a consoleAuthorisation and updates it. Returns the server's representation of the consoleAuthorisation, and an error, if there is any.
func (c *consoleAuthorisations) Update(ctx context.Context, consoleAuthorisation *v1alpha1.ConsoleAuthorisation, opts v1.UpdateOptions) (result *v1alpha1.ConsoleAuthorisation, err error) {
	result = &v1alpha1.ConsoleAuthorisation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("consoleauthorisations").
		Name(consoleAuthorisation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(consoleAuthorisation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *consoleAuthorisations) UpdateStatus(ctx context.Context, consoleAuthorisation *v1alpha1.ConsoleAuthorisation, opts v1.UpdateOptions) (result *v1alpha1.ConsoleAuthorisation, err error) {
	result = &v1alpha1.ConsoleAuthorisation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("consoleauthorisations").
		Name(consoleAuthorisation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(consoleAuthorisation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the consoleAuthorisation and deletes it. Returns an error if one occurs.
func (c *consoleAuthorisations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("consoleauthorisations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *consoleAuthorisations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("consoleauthorisations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched consoleAuthorisation.
func (c *consoleAuthorisations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConsoleAuthorisation, err error) {
	result = &v1alpha1.ConsoleAuthorisation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("consoleauthorisations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

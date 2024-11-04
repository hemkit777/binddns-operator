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

package v1

import (
	"context"
	"time"

	v1 "github.com/bind-dns/binddns-operator/pkg/apis/binddns/v1"
	scheme "github.com/bind-dns/binddns-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DnsDomainsGetter has a method to return a DnsDomainInterface.
// A group's client should implement this interface.
type DnsDomainsGetter interface {
	DnsDomains() DnsDomainInterface
}

// DnsDomainInterface has methods to work with DnsDomain resources.
type DnsDomainInterface interface {
	Create(ctx context.Context, dnsDomain *v1.DnsDomain, opts metav1.CreateOptions) (*v1.DnsDomain, error)
	Update(ctx context.Context, dnsDomain *v1.DnsDomain, opts metav1.UpdateOptions) (*v1.DnsDomain, error)
	UpdateStatus(ctx context.Context, dnsDomain *v1.DnsDomain, opts metav1.UpdateOptions) (*v1.DnsDomain, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DnsDomain, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DnsDomainList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DnsDomain, err error)
	DnsDomainExpansion
}

// dnsDomains implements DnsDomainInterface
type dnsDomains struct {
	client rest.Interface
}

// newDnsDomains returns a DnsDomains
func newDnsDomains(c *BinddnsV1Client) *dnsDomains {
	return &dnsDomains{
		client: c.RESTClient(),
	}
}

// Get takes name of the dnsDomain, and returns the corresponding dnsDomain object, and an error if there is any.
func (c *dnsDomains) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DnsDomain, err error) {
	result = &v1.DnsDomain{}
	err = c.client.Get().
		Resource("dnsdomains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DnsDomains that match those selectors.
func (c *dnsDomains) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DnsDomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DnsDomainList{}
	err = c.client.Get().
		Resource("dnsdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dnsDomains.
func (c *dnsDomains) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("dnsdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dnsDomain and creates it.  Returns the server's representation of the dnsDomain, and an error, if there is any.
func (c *dnsDomains) Create(ctx context.Context, dnsDomain *v1.DnsDomain, opts metav1.CreateOptions) (result *v1.DnsDomain, err error) {
	result = &v1.DnsDomain{}
	err = c.client.Post().
		Resource("dnsdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dnsDomain).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dnsDomain and updates it. Returns the server's representation of the dnsDomain, and an error, if there is any.
func (c *dnsDomains) Update(ctx context.Context, dnsDomain *v1.DnsDomain, opts metav1.UpdateOptions) (result *v1.DnsDomain, err error) {
	result = &v1.DnsDomain{}
	err = c.client.Put().
		Resource("dnsdomains").
		Name(dnsDomain.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dnsDomain).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dnsDomains) UpdateStatus(ctx context.Context, dnsDomain *v1.DnsDomain, opts metav1.UpdateOptions) (result *v1.DnsDomain, err error) {
	result = &v1.DnsDomain{}
	err = c.client.Put().
		Resource("dnsdomains").
		Name(dnsDomain.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dnsDomain).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dnsDomain and deletes it. Returns an error if one occurs.
func (c *dnsDomains) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("dnsdomains").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dnsDomains) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("dnsdomains").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dnsDomain.
func (c *dnsDomains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DnsDomain, err error) {
	result = &v1.DnsDomain{}
	err = c.client.Patch(pt).
		Resource("dnsdomains").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
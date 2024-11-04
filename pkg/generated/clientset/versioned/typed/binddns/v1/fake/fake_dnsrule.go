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
	"context"

	binddnsv1 "github.com/hemkit777/binddns-operator/pkg/apis/binddns/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDnsRules implements DnsRuleInterface
type FakeDnsRules struct {
	Fake *FakeBinddnsV1
}

var dnsrulesResource = schema.GroupVersionResource{Group: "binddns.github.com", Version: "v1", Resource: "dnsrules"}

var dnsrulesKind = schema.GroupVersionKind{Group: "binddns.github.com", Version: "v1", Kind: "DnsRule"}

// Get takes name of the dnsRule, and returns the corresponding dnsRule object, and an error if there is any.
func (c *FakeDnsRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *binddnsv1.DnsRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dnsrulesResource, name), &binddnsv1.DnsRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*binddnsv1.DnsRule), err
}

// List takes label and field selectors, and returns the list of DnsRules that match those selectors.
func (c *FakeDnsRules) List(ctx context.Context, opts v1.ListOptions) (result *binddnsv1.DnsRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dnsrulesResource, dnsrulesKind, opts), &binddnsv1.DnsRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &binddnsv1.DnsRuleList{ListMeta: obj.(*binddnsv1.DnsRuleList).ListMeta}
	for _, item := range obj.(*binddnsv1.DnsRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dnsRules.
func (c *FakeDnsRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dnsrulesResource, opts))
}

// Create takes the representation of a dnsRule and creates it.  Returns the server's representation of the dnsRule, and an error, if there is any.
func (c *FakeDnsRules) Create(ctx context.Context, dnsRule *binddnsv1.DnsRule, opts v1.CreateOptions) (result *binddnsv1.DnsRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dnsrulesResource, dnsRule), &binddnsv1.DnsRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*binddnsv1.DnsRule), err
}

// Update takes the representation of a dnsRule and updates it. Returns the server's representation of the dnsRule, and an error, if there is any.
func (c *FakeDnsRules) Update(ctx context.Context, dnsRule *binddnsv1.DnsRule, opts v1.UpdateOptions) (result *binddnsv1.DnsRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dnsrulesResource, dnsRule), &binddnsv1.DnsRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*binddnsv1.DnsRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDnsRules) UpdateStatus(ctx context.Context, dnsRule *binddnsv1.DnsRule, opts v1.UpdateOptions) (*binddnsv1.DnsRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dnsrulesResource, "status", dnsRule), &binddnsv1.DnsRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*binddnsv1.DnsRule), err
}

// Delete takes name of the dnsRule and deletes it. Returns an error if one occurs.
func (c *FakeDnsRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(dnsrulesResource, name), &binddnsv1.DnsRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDnsRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dnsrulesResource, listOpts)

	_, err := c.Fake.Invokes(action, &binddnsv1.DnsRuleList{})
	return err
}

// Patch applies the patch and returns the patched dnsRule.
func (c *FakeDnsRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *binddnsv1.DnsRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dnsrulesResource, name, pt, data, subresources...), &binddnsv1.DnsRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*binddnsv1.DnsRule), err
}

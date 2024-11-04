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
	v1 "github.com/bind-dns/binddns-operator/pkg/apis/binddns/v1"
	"github.com/bind-dns/binddns-operator/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type BinddnsV1Interface interface {
	RESTClient() rest.Interface
	DnsDomainsGetter
	DnsRulesGetter
}

// BinddnsV1Client is used to interact with features provided by the binddns.github.com group.
type BinddnsV1Client struct {
	restClient rest.Interface
}

func (c *BinddnsV1Client) DnsDomains() DnsDomainInterface {
	return newDnsDomains(c)
}

func (c *BinddnsV1Client) DnsRules() DnsRuleInterface {
	return newDnsRules(c)
}

// NewForConfig creates a new BinddnsV1Client for the given config.
func NewForConfig(c *rest.Config) (*BinddnsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BinddnsV1Client{client}, nil
}

// NewForConfigOrDie creates a new BinddnsV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BinddnsV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BinddnsV1Client for the given RESTClient.
func New(c rest.Interface) *BinddnsV1Client {
	return &BinddnsV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BinddnsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
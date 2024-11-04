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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/bind-dns/binddns-operator/pkg/apis/binddns/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DnsRuleLister helps list DnsRules.
// All objects returned here must be treated as read-only.
type DnsRuleLister interface {
	// List lists all DnsRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DnsRule, err error)
	// Get retrieves the DnsRule from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.DnsRule, error)
	DnsRuleListerExpansion
}

// dnsRuleLister implements the DnsRuleLister interface.
type dnsRuleLister struct {
	indexer cache.Indexer
}

// NewDnsRuleLister returns a new DnsRuleLister.
func NewDnsRuleLister(indexer cache.Indexer) DnsRuleLister {
	return &dnsRuleLister{indexer: indexer}
}

// List lists all DnsRules in the indexer.
func (s *dnsRuleLister) List(selector labels.Selector) (ret []*v1.DnsRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DnsRule))
	})
	return ret, err
}

// Get retrieves the DnsRule from the index for a given name.
func (s *dnsRuleLister) Get(name string) (*v1.DnsRule, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("dnsrule"), name)
	}
	return obj.(*v1.DnsRule), nil
}
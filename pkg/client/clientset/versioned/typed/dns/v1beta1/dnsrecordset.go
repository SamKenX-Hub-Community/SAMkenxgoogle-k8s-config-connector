// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/dns/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DNSRecordSetsGetter has a method to return a DNSRecordSetInterface.
// A group's client should implement this interface.
type DNSRecordSetsGetter interface {
	DNSRecordSets(namespace string) DNSRecordSetInterface
}

// DNSRecordSetInterface has methods to work with DNSRecordSet resources.
type DNSRecordSetInterface interface {
	Create(ctx context.Context, dNSRecordSet *v1beta1.DNSRecordSet, opts v1.CreateOptions) (*v1beta1.DNSRecordSet, error)
	Update(ctx context.Context, dNSRecordSet *v1beta1.DNSRecordSet, opts v1.UpdateOptions) (*v1beta1.DNSRecordSet, error)
	UpdateStatus(ctx context.Context, dNSRecordSet *v1beta1.DNSRecordSet, opts v1.UpdateOptions) (*v1beta1.DNSRecordSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.DNSRecordSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DNSRecordSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DNSRecordSet, err error)
	DNSRecordSetExpansion
}

// dNSRecordSets implements DNSRecordSetInterface
type dNSRecordSets struct {
	client rest.Interface
	ns     string
}

// newDNSRecordSets returns a DNSRecordSets
func newDNSRecordSets(c *DnsV1beta1Client, namespace string) *dNSRecordSets {
	return &dNSRecordSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dNSRecordSet, and returns the corresponding dNSRecordSet object, and an error if there is any.
func (c *dNSRecordSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DNSRecordSet, err error) {
	result = &v1beta1.DNSRecordSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsrecordsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DNSRecordSets that match those selectors.
func (c *dNSRecordSets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DNSRecordSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DNSRecordSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsrecordsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dNSRecordSets.
func (c *dNSRecordSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnsrecordsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dNSRecordSet and creates it.  Returns the server's representation of the dNSRecordSet, and an error, if there is any.
func (c *dNSRecordSets) Create(ctx context.Context, dNSRecordSet *v1beta1.DNSRecordSet, opts v1.CreateOptions) (result *v1beta1.DNSRecordSet, err error) {
	result = &v1beta1.DNSRecordSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnsrecordsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNSRecordSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dNSRecordSet and updates it. Returns the server's representation of the dNSRecordSet, and an error, if there is any.
func (c *dNSRecordSets) Update(ctx context.Context, dNSRecordSet *v1beta1.DNSRecordSet, opts v1.UpdateOptions) (result *v1beta1.DNSRecordSet, err error) {
	result = &v1beta1.DNSRecordSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsrecordsets").
		Name(dNSRecordSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNSRecordSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dNSRecordSets) UpdateStatus(ctx context.Context, dNSRecordSet *v1beta1.DNSRecordSet, opts v1.UpdateOptions) (result *v1beta1.DNSRecordSet, err error) {
	result = &v1beta1.DNSRecordSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsrecordsets").
		Name(dNSRecordSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNSRecordSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dNSRecordSet and deletes it. Returns an error if one occurs.
func (c *dNSRecordSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsrecordsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dNSRecordSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsrecordsets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dNSRecordSet.
func (c *dNSRecordSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DNSRecordSet, err error) {
	result = &v1beta1.DNSRecordSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnsrecordsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
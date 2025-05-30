/*
Copyright Joao Vicente.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NodeStoreSpecApplyConfiguration represents a declarative configuration of the NodeStoreSpec type for use
// with apply.
type NodeStoreSpecApplyConfiguration struct {
	Name      *string                                  `json:"name,omitempty"`
	Selectors map[string]string                        `json:"selectors,omitempty"`
	Tenants   map[string]TenantInfraApplyConfiguration `json:"tenants,omitempty"`
}

// NodeStoreSpecApplyConfiguration constructs a declarative configuration of the NodeStoreSpec type for use with
// apply.
func NodeStoreSpec() *NodeStoreSpecApplyConfiguration {
	return &NodeStoreSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NodeStoreSpecApplyConfiguration) WithName(value string) *NodeStoreSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithSelectors puts the entries into the Selectors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Selectors field,
// overwriting an existing map entries in Selectors field with the same key.
func (b *NodeStoreSpecApplyConfiguration) WithSelectors(entries map[string]string) *NodeStoreSpecApplyConfiguration {
	if b.Selectors == nil && len(entries) > 0 {
		b.Selectors = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Selectors[k] = v
	}
	return b
}

// WithTenants puts the entries into the Tenants field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Tenants field,
// overwriting an existing map entries in Tenants field with the same key.
func (b *NodeStoreSpecApplyConfiguration) WithTenants(entries map[string]TenantInfraApplyConfiguration) *NodeStoreSpecApplyConfiguration {
	if b.Tenants == nil && len(entries) > 0 {
		b.Tenants = make(map[string]TenantInfraApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Tenants[k] = v
	}
	return b
}

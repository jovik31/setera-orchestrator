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
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	seteracomv1 "github/setera/pkg/api/setera.com/v1"

	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// NodeStoreLister helps list NodeStores.
// All objects returned here must be treated as read-only.
type NodeStoreLister interface {
	// List lists all NodeStores in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*seteracomv1.NodeStore, err error)
	// NodeStores returns an object that can list and get NodeStores.
	NodeStores(namespace string) NodeStoreNamespaceLister
	NodeStoreListerExpansion
}

// nodeStoreLister implements the NodeStoreLister interface.
type nodeStoreLister struct {
	listers.ResourceIndexer[*seteracomv1.NodeStore]
}

// NewNodeStoreLister returns a new NodeStoreLister.
func NewNodeStoreLister(indexer cache.Indexer) NodeStoreLister {
	return &nodeStoreLister{listers.New[*seteracomv1.NodeStore](indexer, seteracomv1.Resource("nodestore"))}
}

// NodeStores returns an object that can list and get NodeStores.
func (s *nodeStoreLister) NodeStores(namespace string) NodeStoreNamespaceLister {
	return nodeStoreNamespaceLister{listers.NewNamespaced[*seteracomv1.NodeStore](s.ResourceIndexer, namespace)}
}

// NodeStoreNamespaceLister helps list and get NodeStores.
// All objects returned here must be treated as read-only.
type NodeStoreNamespaceLister interface {
	// List lists all NodeStores in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*seteracomv1.NodeStore, err error)
	// Get retrieves the NodeStore from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*seteracomv1.NodeStore, error)
	NodeStoreNamespaceListerExpansion
}

// nodeStoreNamespaceLister implements the NodeStoreNamespaceLister
// interface.
type nodeStoreNamespaceLister struct {
	listers.ResourceIndexer[*seteracomv1.NodeStore]
}

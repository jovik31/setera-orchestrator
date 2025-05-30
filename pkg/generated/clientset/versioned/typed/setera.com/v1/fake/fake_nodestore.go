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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github/setera/pkg/api/setera.com/v1"
	seteracomv1 "github/setera/pkg/generated/applyconfiguration/setera.com/v1"
	typedseteracomv1 "github/setera/pkg/generated/clientset/versioned/typed/setera.com/v1"

	gentype "k8s.io/client-go/gentype"
)

// fakeNodeStores implements NodeStoreInterface
type fakeNodeStores struct {
	*gentype.FakeClientWithListAndApply[*v1.NodeStore, *v1.NodeStoreList, *seteracomv1.NodeStoreApplyConfiguration]
	Fake *FakeSeteraV1
}

func newFakeNodeStores(fake *FakeSeteraV1, namespace string) typedseteracomv1.NodeStoreInterface {
	return &fakeNodeStores{
		gentype.NewFakeClientWithListAndApply[*v1.NodeStore, *v1.NodeStoreList, *seteracomv1.NodeStoreApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("nodestores"),
			v1.SchemeGroupVersion.WithKind("NodeStore"),
			func() *v1.NodeStore { return &v1.NodeStore{} },
			func() *v1.NodeStoreList { return &v1.NodeStoreList{} },
			func(dst, src *v1.NodeStoreList) { dst.ListMeta = src.ListMeta },
			func(list *v1.NodeStoreList) []*v1.NodeStore { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.NodeStoreList, items []*v1.NodeStore) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}

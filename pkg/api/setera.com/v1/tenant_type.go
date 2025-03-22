// +kubebuilder:object:generate=true
// +groupName=setera.com
// +versionName=v1

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Tenant is a specification for a Tenant resource
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TenantSpec `json:"spec"`
}


type TenantSpec struct {
	Name         string            `json:"name"`      //Tenant Name
	VNI          int               `json:"vni"`       //Tenant VNI identification
	Zones        int               `json:"zones"`     //Number of nodes where the tenant is to be deployed
	Nodes        []Node            `json:"nodes"`     //Node list where the tenant is deployed
	Requirements map[string]string `json:"selectors"` // Tenant requirements in the form of node labels and selectors
}

type Node struct {
	Name    string `json:"name"`              //Node name where the tenant is enabled
	VtepMac string `json:"vtepMac,omitempty"` //VTEP Mac address is saved using string format due to the fact that it generates an error with the cache informer, create string to Mac address
	VtepIp  string `json:"vtepIp,omitempty"`  //IP of the Vtep device for this specific node and tenant
	NodeIP  string `json:"nodeIP,omitempty"`  //Node IP where the tenant is deployed
	Prefix  int    `json:"prefix"`            //Size of tenant CIDR in the node
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TenantList is a list of Tenant resources
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Tenant `json:"items,omitempty"`
}

type ConfMap struct {
	PodCIDR string            `json:"PodCIDR"`
	Backend map[string]string `json:"Backend"`
}

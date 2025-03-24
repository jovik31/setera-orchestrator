// +kubebuilder:rbac:groups=setera.com,resources=tenants,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=setera.com,resources=tenants/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=setera.com,resources=nodeStores,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=setera.com,resources=nodeStores/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch

package agent

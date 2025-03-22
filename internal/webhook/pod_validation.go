package webhook

import (
	"encoding/json"
	"errors"

	"context"
	"log"

	//k8s
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	seterav1 "github/setera/pkg/api/setera.com/v1"
)

/* [ ] Implement pod mutation webhook
   [ ]: Check if pod has the tenant label
   [ ]: If not, add default tenant
   [ ] If it has a tenant validate the tenant
*/

func (ws *WebhookServer) validatePod(admissionRequest *admissionv1.AdmissionRequest) (*admissionv1.AdmissionResponse, error) {

	var pod corev1.Pod
	if err := json.Unmarshal(admissionRequest.Object.Raw, &pod); err != nil {
		log.Printf("Error in unmarshalling pod: %v", err)
	}

	// Fetch tenant list, since tenants are cluster scoped we use "" as the namespace
	tenantList, err := ws.seterav1Clientset.SeteraV1().Tenants("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in fetching tenant list: %v", err)
		return nil, err
	}

	err = checkTenantLabel(pod, tenantList)
	if err != nil {
		log.Printf("Error in checking tenant label: %v", err)
		return nil, err
	}

	return &admissionv1.AdmissionResponse{
		Allowed: true,
		Result: &metav1.Status{
			Message: "pod belongs to an existing tenant",
		},
	}, nil
}

func checkTenantLabel(pod corev1.Pod, tenantList *seterav1.TenantList) error {

	// check if tenant label is present
	tenant_label, ok := pod.Labels["setera.com.v1.tenant"]
	if !ok {
		return errors.New("tenant label not found")
	}

	// check if tenant label is valid
	for _, tenant := range tenantList.Items {
		if tenant_label == tenant.Name {
			return nil
		}
	}

	return errors.New("tenant not found")
}

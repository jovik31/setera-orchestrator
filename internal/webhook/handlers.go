package webhook

import (
	"encoding/json"
	"fmt"

	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func (ws *WebhookServer) validateHandler(w http.ResponseWriter, r *http.Request) {

	//validate http method
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("%s method is not allowed", r.Method), http.StatusMethodNotAllowed)
		klog.Error(http.StatusMethodNotAllowed, fmt.Sprintf(" %s method is not allowed", r.Method))
		return
	}
	//validate headers
	contentType := r.Header.Get(contentTypeHeader)
	if contentType != contentTypeJSON {
		http.Error(w, fmt.Sprintf("%s is not a supported content type", contentType), http.StatusUnsupportedMediaType)
		klog.Error(http.StatusUnsupportedMediaType, fmt.Sprintf("%s is not a supported content type", contentType))
		return
	}
	//check if body is empty
	if r.Body == nil {
		http.Error(w, "request has an empty body", http.StatusBadRequest)
		klog.Error(fmt.Sprintf("error code %d request has an empty body", http.StatusBadRequest))
		return
	}
	//decode request body
	var admissionReview = admissionv1.AdmissionReview{}
	if err := json.NewDecoder(r.Body).Decode(&admissionReview); err != nil {

		http.Error(w, fmt.Sprintf("failed decoding admission request: %s", err), http.StatusBadRequest)
		klog.Error(fmt.Sprintf(" error code: %d failed decoding admission request with error %s", http.StatusBadRequest, err))
		return
	}

	// for struct comparison instead of pointer comparison
	if admissionReview.Request.RequestResource == nil {
		http.Error(w, "failed to extract resource from request", http.StatusBadRequest)
		klog.Error(fmt.Sprintf(" error code: %d failed to extract resource from request", http.StatusBadRequest))
		return
	}
	requestResourceGVK := *admissionReview.Request.RequestResource // check for possible crashout

	//admissionResponse := &admissionv1.AdmissionResponse{}
	switch requestResourceGVK {

	//daemonsets
	case v1.GroupVersionResource{Group: "apps", Version: "v1", Resource: "daemonsets"}:

		w.Write([]byte("This is a daemonset or a deployment"))

	//deployments
	case v1.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}:

		w.Write([]byte("This is a daemonset or a deployment"))

	//pods
	case v1.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}:

		//call pod validation
		admissionResponse, err := ws.validatePod(admissionReview.Request)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)

		}
		w.Write([]byte(fmt.Sprintf("%v", admissionResponse)))

	//tenants
	case v1.GroupVersionResource{Group: "setera.com", Version: "v1", Resource: "tenants"}:

	}

}

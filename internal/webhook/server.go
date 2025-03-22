package webhook

import (
	"net/http"

	seterav1clientset "github/setera/pkg/generated/clientset/versioned"

	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

type WebhookServer struct {
	port string
	//certFile           string
	//keyFile            string
	//tlsDir             string
	server             *http.Server
	seterav1Clientset  seterav1clientset.Interface
	kubernetsClientset kubernetes.Interface

	logger *klog.Logger
}

func NewWebhookServer(seteraClient seterav1clientset.Interface, k8sClientset kubernetes.Interface) *WebhookServer {

	klog.InitFlags(nil)
	return &WebhookServer{
		port: serverPort,
		//certFile:           tlsCertFile,
		//keyFile:            tlsCertKey,
		//tlsDir:             tlsDir,
		seterav1Clientset:  seteraClient,
		kubernetsClientset: k8sClientset,
	}
}

func (ws *WebhookServer) Start() error {

	router := http.NewServeMux()
	router.HandleFunc(validatePodEndpoint, ws.validateHandler)
	//http.HandleFunc(validatePodEndpoint, ws.HandlePodValidation)
	//http.HandleFunc(validateTenantEndpoint, ws.HandleTenantValidation)

	ws.server = &http.Server{
		Addr:    ws.port,
		Handler: router,
	}

	klog.Info("Starting webhook server on ", ws.server.Addr)
	return ws.server.ListenAndServe()
}

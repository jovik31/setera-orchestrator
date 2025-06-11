package webhook

import (
	"net/http"

	seterav1clientset "github/setera/pkg/generated/clientset/versioned"

	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

type WebhookServer struct {
	port               string
	tlsCert            string
	tlsKey             string
	Server             *http.Server
	seterav1Clientset  seterav1clientset.Interface
	kubernetsClientset kubernetes.Interface
}

func NewWebhookServer(seteraClient seterav1clientset.Interface, k8sClientset kubernetes.Interface, tlsCert string, tlsKey string) *WebhookServer {

	return &WebhookServer{
		port:               serverPort,
		tlsCert:            tlsCert,
		tlsKey:             tlsKey,
		seterav1Clientset:  seteraClient,
		kubernetsClientset: k8sClientset,
	}
}

func (ws *WebhookServer) Start() error {

	router := http.NewServeMux()
	router.HandleFunc(validateEndpoint, ws.admissionValidationHandler)

	middleware := runMiddleware(
		loggingMiddleware,
		validatingMiddleware,
	)

	ws.Server = &http.Server{
		Addr:    ws.port,
		Handler: middleware(router),
	}

	klog.Info("started webhook server at", ws.Server.Addr)
	return ws.Server.ListenAndServeTLS(ws.tlsCert, ws.tlsKey)
}

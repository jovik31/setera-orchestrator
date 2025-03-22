package main

import (
	"github/setera/internal/webhook"
	"github/setera/pkg/k8s"
	"log"
)

/*
	[ ] Implement webhook process

[ ] Implement the webhook service
[ ] Implement the webhook service account
[ ] Implement the webhook rbac
[ ] Implement the webhook deployment
[ ] Implement the webhook certificates
*/
func main() {

	config, err := k8s.InitKubeConfig()
	if err != nil {
		log.Fatal(err, "Error in building kubeconfig")
	}

	sClientset, err := k8s.NewSeteraClient(config)
	if err != nil {
		log.Fatal(err, "Error in building setera clientset")
	}

	kubeClientset, err := k8s.NewKubeClient(config)
	if err != nil {
		log.Fatal(err, "Error in building kubernetes clientset")
	}

	ws := webhook.NewWebhookServer(sClientset, kubeClientset)

	ws.Start()

}

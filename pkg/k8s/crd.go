package k8s

import (
	seterav1clientset "github/setera/pkg/generated/clientset/versioned"
	"log"

	"k8s.io/client-go/rest"
)

func NewSeteraClient(c *rest.Config) (seterav1clientset.Interface, error) {
	// [ ]: Create a new client

	sClientset, err := seterav1clientset.NewForConfig(c)
	if err != nil {
		log.Printf("Error in building setera clientset: %v", err)
		return nil, err
	}
	return sClientset, nil
}

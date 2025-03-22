package k8s

import (
	"flag"
	"log"
	"path/filepath"

	// k8s client-go
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// [ ]: Implement kubeconfig intialization

func InitKubeConfig() (*rest.Config, error) {

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {

		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Printf("Error in building kubeconfig from flags: %v, using in cluster config", err)
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Printf("Error in building in cluster config: %v", err)
			return nil, err
		}
	}

	return config, nil
}

func NewKubeClient(c *rest.Config) (*kubernetes.Clientset, error) {
	// [ ]: Create a new client

	kubeClientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		log.Printf("Error in building kubernetes clientset: %v", err)
		return nil, err
	}
	return kubeClientset, nil
}

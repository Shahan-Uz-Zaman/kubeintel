package kubernetes

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var Clientset *kubernetes.Clientset

// NewClient initializes the Kubernetes client and stores it globally.
func NewClient() (*kubernetes.Clientset, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	kubeconfig := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	Clientset = clientset
	return Clientset, nil
}

// GetClient returns the initialized Kubernetes client.
// If the client has not been initialized yet, it creates one.
func GetClient() (*kubernetes.Clientset, error) {
	if Clientset != nil {
		return Clientset, nil
	}

	return NewClient()
}

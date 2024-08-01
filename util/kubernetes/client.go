package util

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientInstance Client
)

type Client struct {
	Clientset kubernetes.Interface
}

func GetK8sClient() Client {
	if clientInstance.Clientset == nil {
		var kubeconfig string = os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = os.Getenv("HOME") + "/.kube/config"
		}

		// use the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			// attempt to use in cluster if failed to get kubeconfig
			config, err = rest.InClusterConfig()
			if err != nil {
				panic(err.Error())
			}
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		clientInstance = Client{
			Clientset: clientset,
		}
	}
	return clientInstance
}

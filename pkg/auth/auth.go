package auth

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func GetClients(conf string) (*kubernetes.Clientset, metricsv.Interface) {
	config, _ := clientcmd.BuildConfigFromFlags("", conf)
	resourceClient, _ := kubernetes.NewForConfig(config)
	metricsClient, _ := metricsv.NewForConfig(config)

	return resourceClient, metricsClient
}

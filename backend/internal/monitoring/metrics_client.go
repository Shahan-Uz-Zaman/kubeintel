package monitoring

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metricsclient "k8s.io/metrics/pkg/client/clientset/versioned"
)

func NewMetricsClient() (*metricsclient.Clientset, error) {

	config, err := rest.InClusterConfig()

	if err != nil {

		config, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

		if err != nil {
			return nil, err
		}
	}

	return metricsclient.NewForConfig(config)
}
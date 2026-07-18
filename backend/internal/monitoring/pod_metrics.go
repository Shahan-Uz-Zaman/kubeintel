package monitoring

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodMetrics() ([]PodMetric, error) {

	client, err := NewMetricsClient()
	if err != nil {
		return nil, err
	}

	podMetrics, err := client.
		MetricsV1beta1().
		PodMetricses("").
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var result []PodMetric

	for _, pod := range podMetrics.Items {

		var totalCPU int64
		var totalMemory int64

		for _, container := range pod.Containers {

			totalCPU += container.Usage.Cpu().MilliValue()
			totalMemory += container.Usage.Memory().Value()
		}

		result = append(result, PodMetric{

			Namespace: pod.Namespace,

			Name: pod.Name,

			CPU: pod.Containers[0].Usage.Cpu().String(),

			Memory: pod.Containers[0].Usage.Memory().String(),

			CPUUsage: totalCPU,

			MemUsage: totalMemory / (1024 * 1024),
		})
	}

	return result, nil
}
package monitoring

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeMetrics() ([]NodeMetric, error) {

	client, err := NewMetricsClient()
	if err != nil {
		return nil, err
	}

	nodeMetrics, err := client.MetricsV1beta1().
		NodeMetricses().
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var result []NodeMetric

	for _, node := range nodeMetrics.Items {

		cpu := node.Usage.Cpu()
		mem := node.Usage.Memory()

		result = append(result, NodeMetric{
			Name:     node.Name,
			CPU:      cpu.String(),
			Memory:   mem.String(),
			CPUUsage: cpu.MilliValue(),
			MemUsage: mem.Value() / (1024 * 1024), // MB
		})
	}

	return result, nil
}

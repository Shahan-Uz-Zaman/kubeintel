package monitoring

type NodeMetric struct {
	Name      string `json:"name"`
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
	CPUUsage  int64  `json:"cpuUsage"`
	MemUsage  int64  `json:"memoryUsage"`
}

type PodMetric struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
	CPUUsage  int64  `json:"cpuUsage"`
	MemUsage  int64  `json:"memoryUsage"`
}
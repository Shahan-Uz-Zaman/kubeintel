package models

type DashboardResponse struct {
	ClusterStatus string `json:"clusterStatus"`

	NodeCount int `json:"nodeCount"`

	NamespaceCount int `json:"namespaceCount"`

	RunningPods int `json:"runningPods"`

	FailedPods int `json:"failedPods"`
}
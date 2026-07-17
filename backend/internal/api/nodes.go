package api

import (
	"context"
	"net/http"

	"kubeintel/backend/internal/kubernetes"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NodeResponse struct {
	Name              string `json:"name"`
	Status            string `json:"status"`
	Roles             string `json:"roles"`
	KubernetesVersion string `json:"kubernetesVersion"`
	OS                string `json:"os"`
	Architecture      string `json:"architecture"`
	ContainerRuntime  string `json:"containerRuntime"`
	InternalIP        string `json:"internalIP"`
	ExternalIP        string `json:"externalIP"`
	CreationTime      string `json:"creationTime"`
}

func GetNodes(c *gin.Context) {

	client := kubernetes.Clientset

	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Kubernetes client not initialized",
		})
		return
	}

	nodes, err := client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var response []NodeResponse

	for _, node := range nodes.Items {

		status := "Unknown"

		for _, condition := range node.Status.Conditions {
			if condition.Type == "Ready" {
				if condition.Status == "True" {
					status = "Ready"
				} else {
					status = "NotReady"
				}
			}
		}

		internalIP := ""
		externalIP := ""

		for _, address := range node.Status.Addresses {

			if address.Type == "InternalIP" {
				internalIP = address.Address
			}

			if address.Type == "ExternalIP" {
				externalIP = address.Address
			}
		}

		roles := "worker"

		if _, ok := node.Labels["node-role.kubernetes.io/control-plane"]; ok {
			roles = "control-plane"
		}

		if _, ok := node.Labels["node-role.kubernetes.io/master"]; ok {
			roles = "master"
		}

		response = append(response, NodeResponse{
			Name:              node.Name,
			Status:            status,
			Roles:             roles,
			KubernetesVersion: node.Status.NodeInfo.KubeletVersion,
			OS:                node.Status.NodeInfo.OperatingSystem,
			Architecture:      node.Status.NodeInfo.Architecture,
			ContainerRuntime:  node.Status.NodeInfo.ContainerRuntimeVersion,
			InternalIP:        internalIP,
			ExternalIP:        externalIP,
			CreationTime:      node.CreationTimestamp.String(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(response),
		"nodes": response,
	})
}
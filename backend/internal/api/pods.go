package api

import (
	"context"
	"net/http"

	"kubeintel/backend/internal/kubernetes"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodResponse struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Node         string `json:"node"`
	Status       string `json:"status"`
	PodIP        string `json:"podIP"`
	HostIP       string `json:"hostIP"`
	Restarts     int32  `json:"restarts"`
	CreationTime string `json:"creationTime"`
}

func GetPods(c *gin.Context) {

	client := kubernetes.Clientset

	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Kubernetes client not initialized",
		})
		return
	}

	pods, err := client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var response []PodResponse

	for _, pod := range pods.Items {

		var restarts int32

		for _, container := range pod.Status.ContainerStatuses {
			restarts += container.RestartCount
		}

		response = append(response, PodResponse{
			Name:         pod.Name,
			Namespace:    pod.Namespace,
			Node:         pod.Spec.NodeName,
			Status:       string(pod.Status.Phase),
			PodIP:        pod.Status.PodIP,
			HostIP:       pod.Status.HostIP,
			Restarts:     restarts,
			CreationTime: pod.CreationTimestamp.String(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(response),
		"pods":  response,
	})
}

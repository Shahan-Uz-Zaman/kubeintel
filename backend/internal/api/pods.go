package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeintel/backend/internal/kubernetes"
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

func getPodHealthStatus(pod corev1.Pod) (string, string, int32) {

	var totalRestarts int32
	status := "Healthy"
	reason := ""

	for _, cs := range pod.Status.ContainerStatuses {

		totalRestarts += cs.RestartCount

		if cs.State.Waiting != nil {

			reason = cs.State.Waiting.Reason

			switch reason {

			case "CrashLoopBackOff":
				status = "Critical"

			case "ImagePullBackOff":
				status = "Critical"

			case "ErrImagePull":
				status = "Critical"

			default:
				status = "Warning"
			}
		}

		if cs.LastTerminationState.Terminated != nil {

			if cs.LastTerminationState.Terminated.Reason == "OOMKilled" {

				status = "Critical"
				reason = "OOMKilled"
			}
		}

		if cs.RestartCount > 5 && status == "Healthy" {

			status = "Warning"
			reason = "High Restart Count"
		}
	}

	if pod.Status.Phase == corev1.PodPending {

		status = "Warning"
		reason = "Pending"
	}

	if pod.Status.Phase == corev1.PodFailed {

		status = "Critical"
		reason = "Failed"
	}

	return status, reason, totalRestarts
}

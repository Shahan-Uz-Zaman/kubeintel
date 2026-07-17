package api

import (
	"context"
	"net/http"

	"kubeintel/backend/internal/kubernetes"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetCluster(c *gin.Context) {

	client := kubernetes.Clientset

	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Kubernetes client not initialized",
		})
		return
	}

	version, err := client.Discovery().ServerVersion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	nodes, _ := client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	namespaces, _ := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})

	pods, _ := client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})

	var runningPods int
	var pendingPods int
	var failedPods int

	for _, pod := range pods.Items {

		switch string(pod.Status.Phase) {

		case "Running":
			runningPods++

		case "Pending":
			pendingPods++

		case "Failed":
			failedPods++
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"clusterVersion": version.GitVersion,
		"platform":       version.Platform,
		"nodes":          len(nodes.Items),
		"namespaces":     len(namespaces.Items),
		"pods":           len(pods.Items),
		"runningPods":    runningPods,
		"pendingPods":    pendingPods,
		"failedPods":     failedPods,
	})
}
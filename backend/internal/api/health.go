package api

import (
	"context"
	"net/http"

	"kubeintel/backend/internal/kubernetes"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetHealth(c *gin.Context) {

	client := kubernetes.Clientset

	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Kubernetes client not initialized",
		})
		return
	}

	pods, err := client.CoreV1().
		Pods("").
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	nodes, err := client.CoreV1().
		Nodes().
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var runningPods int
	var pendingPods int
	var failedPods int

	for _, pod := range pods.Items {

		switch pod.Status.Phase {

		case "Running":
			runningPods++

		case "Pending":
			pendingPods++

		case "Failed":
			failedPods++
		}
	}

	var healthyNodes int
	var unhealthyNodes int

	for _, node := range nodes.Items {

		ready := false

		for _, condition := range node.Status.Conditions {

			if condition.Type == "Ready" {

				if condition.Status == "True" {
					ready = true
				}

				break
			}
		}

		if ready {
			healthyNodes++
		} else {
			unhealthyNodes++
		}
	}

	score := 100

	score -= failedPods * 15
	score -= pendingPods * 5
	score -= unhealthyNodes * 20

	if score < 0 {
		score = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"totalPods":      len(pods.Items),
		"runningPods":    runningPods,
		"pendingPods":    pendingPods,
		"failedPods":     failedPods,
		"totalNodes":     len(nodes.Items),
		"healthyNodes":   healthyNodes,
		"unhealthyNodes": unhealthyNodes,
		"healthScore":    score,
	})
}

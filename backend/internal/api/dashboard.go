package api

import (
	"net/http"

	"kubeintel/backend/internal/models"
	"kubeintel/backend/internal/kubernetes"

	corev1 "k8s.io/api/core/v1"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDashboard(c *gin.Context) {

	clientset := kubernetes.Clientset

	if clientset == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Kubernetes client not initialized",
		})
		return
	}

	// -------------------------
	// Get Nodes
	// -------------------------

	nodes, err := clientset.CoreV1().Nodes().List(c, metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// -------------------------
	// Get Namespaces
	// -------------------------

	namespaces, err := clientset.CoreV1().Namespaces().List(c, metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// -------------------------
	// Get Pods
	// -------------------------

	pods, err := clientset.CoreV1().Pods("").List(c, metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	runningPods := 0
	failedPods := 0

	clusterHealthy := true

	for _, pod := range pods.Items {

		switch pod.Status.Phase {

		case corev1.PodRunning:
			runningPods++

		case corev1.PodFailed:
			failedPods++
			clusterHealthy = false
		}
	}

	clusterStatus := "Healthy"

	if !clusterHealthy {
		clusterStatus = "Warning"
	}

	response := models.DashboardResponse{
		ClusterStatus: clusterStatus,

		NodeCount: len(nodes.Items),

		NamespaceCount: len(namespaces.Items),

		RunningPods: runningPods,

		FailedPods: failedPods,
	}

	c.JSON(http.StatusOK, response)
}
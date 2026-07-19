package api

import (
	"context"
	"io"

	"github.com/gin-gonic/gin"

	corev1 "k8s.io/api/core/v1"

	"kubeintel/backend/internal/kubernetes"
)

func GetPodLogs(c *gin.Context) {

	namespace := c.DefaultQuery("namespace", "default")
	pod := c.Query("pod")

	if pod == "" {
		c.JSON(400, gin.H{
			"error": "pod is required",
		})
		return
	}

	client, err := kubernetes.NewClient()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	req := client.CoreV1().
		Pods(namespace).
		GetLogs(pod, &corev1.PodLogOptions{
			TailLines: func() *int64 {
				t := int64(200)
				return &t
			}(),
		})

	stream, err := req.Stream(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer stream.Close()

	bytes, _ := io.ReadAll(stream)

	c.JSON(200, gin.H{
		"namespace": namespace,
		"pod":       pod,
		"logs":      string(bytes),
	})
}

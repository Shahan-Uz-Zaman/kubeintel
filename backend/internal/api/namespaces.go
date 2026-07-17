package api

import (
	"context"
	"net/http"

	"kubeintel/backend/internal/kubernetes"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceResponse struct {
	Name         string `json:"name"`
	Status       string `json:"status"`
	CreationTime string `json:"creationTime"`
}

func GetNamespaces(c *gin.Context) {

	client := kubernetes.Clientset

	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Kubernetes client not initialized",
		})
		return
	}

	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var response []NamespaceResponse

	for _, ns := range namespaces.Items {

		response = append(response, NamespaceResponse{
			Name:         ns.Name,
			Status:       string(ns.Status.Phase),
			CreationTime: ns.CreationTimestamp.String(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"count":      len(response),
		"namespaces": response,
	})
}
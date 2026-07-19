package api

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kubeintel/backend/internal/kubernetes"
)

type EventResponse struct {
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	Object    string `json:"object"`
	Namespace string `json:"namespace"`
	Message   string `json:"message"`
	Time      string `json:"time"`
}

func GetEvents(c *gin.Context) {

	namespace := c.DefaultQuery("namespace", "default")

	client, err := kubernetes.NewClient()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	events, err := client.CoreV1().
		Events(namespace).
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := []EventResponse{}

	for _, e := range events.Items {

		eventTime := ""

		if !e.EventTime.Time.IsZero() {
			eventTime = e.EventTime.Time.Format(time.RFC3339)
		} else if !e.LastTimestamp.IsZero() {
			eventTime = e.LastTimestamp.Time.Format(time.RFC3339)
		} else {
			eventTime = e.CreationTimestamp.Time.Format(time.RFC3339)
		}

		response = append(response, EventResponse{
			Type:      e.Type,
			Reason:    e.Reason,
			Object:    e.InvolvedObject.Name,
			Namespace: e.Namespace,
			Message:   e.Message,
			Time:      eventTime,
		})
	}

	c.JSON(200, gin.H{
		"events": response,
	})
}

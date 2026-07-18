package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kubeintel/backend/internal/monitoring"
)

func GetNodeMetrics(c *gin.Context) {

	nodes, err := monitoring.GetNodeMetrics()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nodes)
}

func GetPodMetrics(c *gin.Context) {

	pods, err := monitoring.GetPodMetrics()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pods)
}

func GetClusterMetrics(c *gin.Context) {

	nodes, err := monitoring.GetNodeMetrics()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	pods, err := monitoring.GetPodMetrics()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := gin.H{
		"nodes": nodes,
		"pods":  pods,
	}

	c.JSON(http.StatusOK, response)
}
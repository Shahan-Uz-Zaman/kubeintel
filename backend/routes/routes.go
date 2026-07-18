package routes

import (
	"github.com/gin-gonic/gin"

	"kubeintel/backend/internal/api"
)

func SetupRoutes(router *gin.Engine) {

	// Health
	router.GET("/health", api.Health)

	// Assignment 2
	router.GET("/api/cluster", api.GetCluster)
	router.GET("/api/nodes", api.GetNodes)
	router.GET("/api/pods", api.GetPods)
	router.GET("/api/namespaces", api.GetNamespaces)
	

	// Assignment 3
	router.GET("/api/dashboard", api.GetDashboard)

	// Assignment 4
	monitoring := router.Group("/api/monitoring")
	{
		monitoring.GET("/nodes", api.GetNodeMetrics)
		monitoring.GET("/pods", api.GetPodMetrics)
		monitoring.GET("/cluster", api.GetClusterMetrics)
		monitoring.GET("/network", api.GetNetwork)
		monitoring.GET("/storage", api.GetStorage)
	}
}
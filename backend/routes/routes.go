package routes

import (
	"github.com/gin-gonic/gin"

	"kubeintel/backend/internal/api"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/health", api.Health)

	router.GET("/api/cluster", api.GetCluster)

	router.GET("/api/nodes", api.GetNodes)

	router.GET("/api/pods", api.GetPods)

	router.GET("/api/namespaces", api.GetNamespaces)

	// Assignment 3
	router.GET("/api/dashboard", api.GetDashboard)
}
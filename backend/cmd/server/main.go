package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"time"

	"github.com/gin-contrib/cors"
	"kubeintel/backend/internal/kubernetes"
	"kubeintel/backend/routes"
)

func main() {

	client, err := kubernetes.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	version, err := client.Discovery().ServerVersion()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Kubernetes:", version.GitVersion)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.SetTrustedProxies(nil)

	routes.SetupRoutes(router)

	log.Println("Server running on :8080")

	router.Run(":8080")
}

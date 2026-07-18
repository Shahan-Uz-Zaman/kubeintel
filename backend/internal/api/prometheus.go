package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kubeintel/backend/internal/prometheus"
)

func GetStorage(c *gin.Context) {

	data, err := prometheus.GetStorage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetNetwork(c *gin.Context) {

	data, err := prometheus.GetNetwork()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
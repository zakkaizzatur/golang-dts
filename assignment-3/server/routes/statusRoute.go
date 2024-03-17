package routes

import (
	"assignment-3/controllers"
	"assignment-3/models"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupStatusRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // Update with your frontend URL
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
	
	router.GET("/status", getStatus)
}

func getStatus(c *gin.Context) {
	status := controllers.GetStatus()

	response := models.Response{
		Status: status,
	}

	c.JSON(http.StatusOK, response)
}
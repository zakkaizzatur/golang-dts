package main

import (
	"assignment-3/controllers"
	"assignment-3/routes"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	router := gin.Default()

	routes.SetupStatusRoutes(router)

	go func() {
		for {
			controllers.UpdateStatus()

			time.Sleep(15 * time.Second)
		}
	}()

	router.Run(":8080")
}
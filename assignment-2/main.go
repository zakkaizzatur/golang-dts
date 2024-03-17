package main

import (
	database "assignment-2/config"
	"assignment-2/repositories"
	"assignment-2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := database.StartDB()

	orderRepo := &repositories.OrderRepository{DB:db}

	router := gin.Default()
	routes.SetupOrderRoutes(router, orderRepo)

	router.Run(":8089")
}
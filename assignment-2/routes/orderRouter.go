package routes

import (
	"assignment-2/controller"
	"assignment-2/repositories"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderRepo *repositories.OrderRepository) {
	orderController := &controller.OrderController{OrderRepo: orderRepo,}

	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders", orderController.Get)
	router.GET("/orders/:id", orderController.GetOrderById)
	router.PUT("/orders/:id", orderController.Put)
	router.DELETE("/orders/:id", orderController.Delete)
}
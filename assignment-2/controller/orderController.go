package controller

import (
	"assignment-2/models"
	"assignment-2/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderRepo *repositories.OrderRepository
}

func (h *OrderController) CreateOrder(c *gin.Context) {
	var order models.Orders

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.OrderRepo.Create(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": order,
	})
}

func (h *OrderController) GetOrderById(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.OrderRepo.GetById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (h *OrderController) Get(c *gin.Context) {
	orders, err := h.OrderRepo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (h *OrderController) Delete(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := h.OrderRepo.Delete(uint(orderID)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func (h *OrderController) Put(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)

	existingOrder, err := h.OrderRepo.GetById(uint(orderID))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var updatedOrder models.Orders
    if err := c.ShouldBindJSON(&updatedOrder); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

	*existingOrder = updatedOrder

	if err := h.OrderRepo.UpdateByID(uint(orderID), *existingOrder); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
        return
    }

    // Return success response
    c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "order": existingOrder})
}




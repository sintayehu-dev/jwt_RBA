package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}






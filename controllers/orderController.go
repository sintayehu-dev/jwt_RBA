package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}	

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}






package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)	

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}






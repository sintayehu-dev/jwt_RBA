package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}








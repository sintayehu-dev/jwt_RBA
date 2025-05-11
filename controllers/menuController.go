package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello, World!")
	}
}





package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sintayehu-dev/go_jwt_auth/databases"
	"github.com/sintayehu-dev/go_jwt_auth/middleware"
	"github.com/sintayehu-dev/go_jwt_auth/models"
	routes "github.com/sintayehu-dev/go_jwt_auth/routes"
)

func main() {

	db := databases.InitDB()
	db.AutoMigrate(&models.User{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	router.Use(middleware.Authenticate())

	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}

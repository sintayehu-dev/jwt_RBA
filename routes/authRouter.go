package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sintayehu-dev/go_jwt_auth/controllers"

)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())

	// Auth routes will be defined here
}

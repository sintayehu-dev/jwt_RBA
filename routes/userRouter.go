package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sintayehu-dev/go_jwt_auth/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}

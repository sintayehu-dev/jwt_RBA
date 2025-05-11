package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sintayehu-dev/go_jwt_auth/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}


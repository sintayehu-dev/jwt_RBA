package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sintayehu-dev/go_jwt_auth/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}

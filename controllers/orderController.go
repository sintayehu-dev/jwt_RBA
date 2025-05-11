package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sintayehu-dev/go_jwt_auth/databases"
	"github.com/sintayehu-dev/go_jwt_auth/helpers"
	"github.com/sintayehu-dev/go_jwt_auth/models"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var orders []models.Order
		if err := databases.DB.WithContext(ctx).Find(&orders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve orders. Please try again later."})
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this order"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderId := c.Param("order_id")
		var order models.Order

		if err := databases.DB.WithContext(ctx).Where("order_id = ?", orderId).First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "The requested order could not be found"})
			return
		}

		if err := helpers.MatchUserTypeToUid(c, order.UserID); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this order"})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to create an order"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := databases.DB.WithContext(ctx).Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create order. Please try again later."})
			return
		}

		c.JSON(http.StatusCreated, order)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update an order"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderId := c.Param("order_id")
		var order models.Order

		if err := databases.DB.WithContext(ctx).Where("order_id = ?", orderId).First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "The requested order could not be found"})
			return
		}

		var updateData models.Order
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if updateData.OrderID != orderId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The order ID in the request body does not match the URL"})
			return
		}

		if err := databases.DB.WithContext(ctx).Where("order_id = ?", orderId).Updates(&updateData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update order. Please try again later."})
			return
		}

		if err := databases.DB.WithContext(ctx).Where("order_id = ?", orderId).First(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve updated order. Please try again later."})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete an order"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderId := c.Param("order_id")

		// Check if order exists
		var order models.Order
		if err := databases.DB.WithContext(ctx).Where("order_id = ?", orderId).First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "The requested order could not be found"})
			return
		}

		// Check for related invoice before deleting
		var invoiceCount int64
		if err := databases.DB.WithContext(ctx).Model(&models.Invoice{}).Where("order_id = ?", orderId).Count(&invoiceCount).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check related invoices. Please try again later."})
			return
		}

		if invoiceCount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "This order cannot be deleted because it has associated invoices"})
			return
		}

		// Delete associated order items first
		if err := databases.DB.WithContext(ctx).Where("order_id = ?", orderId).Delete(&models.OrderItem{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete related order items. Please try again later."})
			return
		}

		// Delete the order
		if err := databases.DB.WithContext(ctx).Where("order_id = ?", orderId).Delete(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete order. Please try again later."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order and associated items deleted successfully"})
	}
}

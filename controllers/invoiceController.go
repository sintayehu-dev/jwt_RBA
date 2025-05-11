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

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoices []models.Invoice
		if err := databases.DB.WithContext(ctx).Find(&invoices).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching invoices"})
			return
		}

		c.JSON(http.StatusOK, invoices)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice

		if err := databases.DB.WithContext(ctx).Where("invoice_id = ?", invoiceId).First(&invoice).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
			return
		}

		var order models.Order
		if err := databases.DB.WithContext(ctx).Where("order_id = ?", invoice.OrderID).First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "related order not found"})
			return
		}

		if err := helpers.MatchUserTypeToUid(c, order.UserID); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, invoice)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice models.Invoice
		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var orderExists int64
		if err := databases.DB.WithContext(ctx).Model(&models.Order{}).Where("order_id = ?", invoice.OrderID).Count(&orderExists).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking order"})
			return
		}

		if orderExists == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
			return
		}

		if err := databases.DB.WithContext(ctx).Create(&invoice).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating invoice"})
			return
		}

		c.JSON(http.StatusCreated, invoice)
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice

		if err := databases.DB.WithContext(ctx).Where("invoice_id = ?", invoiceId).First(&invoice).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
			return
		}

		var updateData models.Invoice
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if updateData.InvoiceID != invoiceId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invoice_id in path and body must match"})
			return
		}

		if err := databases.DB.WithContext(ctx).Where("invoice_id = ?", invoiceId).Updates(&updateData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating invoice"})
			return
		}

		if err := databases.DB.WithContext(ctx).Where("invoice_id = ?", invoiceId).First(&invoice).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching updated invoice"})
			return
		}

		c.JSON(http.StatusOK, invoice)
	}
}


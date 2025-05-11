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

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menus []models.Menu
		if err := databases.DB.WithContext(ctx).Find(&menus).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching menus"})
			return
		}

		c.JSON(http.StatusOK, menus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		var menu models.Menu

		if err := databases.DB.WithContext(ctx).Where("menu_id = ?", menuId).First(&menu).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "menu not found"})
			return	
		}

		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu
		if err := c.ShouldBindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := databases.DB.WithContext(ctx).Create(&menu).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating menu"})
			return
		}

		c.JSON(http.StatusCreated, menu)

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		var menu models.Menu

		if err := databases.DB.WithContext(ctx).Where("menu_id = ?", menuId).First(&menu).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "menu not found"})
			return
		}

		var updateData models.Menu
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if updateData.MenuID != menuId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "menu_id in path and body must match"})					
			return
		}

		if err := databases.DB.WithContext(ctx).Where("menu_id = ?", menuId).Updates(&updateData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating menu"})
			return
		}

		if err := databases.DB.WithContext(ctx).Where("menu_id = ?", menuId).First(&menu).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching updated menu"})
			return
		}

		c.JSON(http.StatusOK, menu)
	}
}
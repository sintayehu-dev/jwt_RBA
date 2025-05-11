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
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view all menus"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menus []models.Menu
		if err := databases.DB.WithContext(ctx).Find(&menus).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve menus. Please try again later."})
			return
		}

		c.JSON(http.StatusOK, menus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this menu"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		var menu models.Menu

		if err := databases.DB.WithContext(ctx).Where("menu_id = ?", menuId).First(&menu).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "The requested menu could not be found"})
			return
		}

		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to create menus"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu
		if err := c.ShouldBindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu data provided. Please check your input."})
			return
		}

		if err := databases.DB.WithContext(ctx).Create(&menu).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create menu. Please try again later."})
			return
		}

		c.JSON(http.StatusCreated, menu)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update menus"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		var menu models.Menu

		if err := databases.DB.WithContext(ctx).Where("menu_id = ?", menuId).First(&menu).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "The menu you're trying to update could not be found"})
			return
		}

		var updateData models.Menu
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu data provided. Please check your input."})
			return
		}

		if updateData.MenuID != menuId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The menu ID in the request does not match the URL"})
			return
		}

		if err := databases.DB.WithContext(ctx).Where("menu_id = ?", menuId).Updates(&updateData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update menu. Please try again later."})
			return
		}

		if err := databases.DB.WithContext(ctx).Where("menu_id = ?", menuId).First(&menu).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu was updated but could not be retrieved"})
			return
		}

		c.JSON(http.StatusOK, menu)
	}
}

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

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var foods []models.Food
		if err := databases.DB.Find(&foods).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, foods)
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		var food models.Food
		err := databases.DB.WithContext(ctx).Where("food_id = ?", foodId).First(&food).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching the food by id"})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := databases.DB.WithContext(ctx).Create(&food).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while creating the food item"})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := databases.DB.WithContext(ctx).Save(&food).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating the food item"})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		result := databases.DB.WithContext(ctx).Where("food_id = ?", foodId).Delete(&models.Food{})

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting the food item"})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "food item not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
	}
}

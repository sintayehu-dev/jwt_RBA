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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve food items. Please try again later."})
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
			c.JSON(http.StatusNotFound, gin.H{"error": "The requested food item could not be found"})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to add food items to the menu"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food data provided. Please check your input."})
			return
		}
		err := databases.DB.WithContext(ctx).Create(&food).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add food item to the menu. Please try again later."})
			return
		}
		c.JSON(http.StatusCreated, food)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update food items"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		var existingFood models.Food

		if err := databases.DB.WithContext(ctx).Where("food_id = ?", foodId).First(&existingFood).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "The food item you're trying to update could not be found"})
			return
		}

		var food models.Food
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food data provided. Please check your input."})
			return
		}

		if food.FoodID != foodId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The food ID in the request does not match the URL"})
			return
		}

		err := databases.DB.WithContext(ctx).Save(&food).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update food item. Please try again later."})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to remove food items from the menu"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		result := databases.DB.WithContext(ctx).Where("food_id = ?", foodId).Delete(&models.Food{})

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove food item. Please try again later."})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "The food item you're trying to delete could not be found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Food item has been successfully removed from the menu"})
	}
}

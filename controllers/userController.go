package controllers

import (
	"context"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/sintayehu-dev/go_jwt_auth/helpers"
	models "github.com/sintayehu-dev/go_jwt_auth/models"
	"github.com/go-playground/validator/v10"
	databases "github.com/sintayehu-dev/go_jwt_auth/databases"
	
)

var userCollection *mongo.Collection = databases.OpenCollection(databases.Client, "users")
var validate = validator.New()

func HashPassword()
func VerifyPassord()
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the email"})
			return
		}
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
	}													
	 user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	 user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	 user.ID = primitive.NewObjectID()
	 user.User_id = user.ID.Hex()
	 token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_Name, *user.Last_Name, *user.User_Type, user.User_id)
	 user.Token = &token
	 user.Refresh_token = &refreshToken
	 resultInsertNumber, err := userCollection.InsertOne(ctx, user)
	 if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user item was not created"})
		return
	 }
	 defer cancel()
	 c.JSON(http.StatusOK, resultInsertNumber)
	 

}}
func Login()                                                    
func GetUsers()                                                 
func GetUser()
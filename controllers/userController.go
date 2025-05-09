package controllers

import (
	"github.com/go-playground/validator/v10"
	databases "github.com/sintayehu-dev/go_jwt_auth/databases"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = databases.OpenCollection(databases.Client, "users")
var validate = validator.New()

func HashPassword()
func VerifyPassord()
func Signup()                                             
func Login()                                                    
func GetUsers()                                                 
func GetUser()                                               

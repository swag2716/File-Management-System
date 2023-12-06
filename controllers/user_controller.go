package controllers

import (
	"File-Management-System/database"
	"File-Management-System/helpers"
	"File-Management-System/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.LogInUser
		var foundUser models.SignUpUser

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)

		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// if foundUser.Email == nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		// 	return
		// }

		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.Name, foundUser.User_id)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		foundUser.Password = nil

		c.JSON(http.StatusOK, foundUser)
	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// it extracts the JSON data from the request body and populate the user struct with the corresponding field values
		var user models.SignUpUser
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// it validates if the structs fields satisfy the validation tags defined on them
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// to check if the same email already exist
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error ocurred while checking for the email"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Email already exist"})
			return
		}

		// to hash password and store it
		password := HashPassword(*user.Password)
		user.Password = &password

		// to check if the same phone number already exists
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error ocurred while checking for phone number"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Phone number already exist"})
			return
		}

		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.Name, user.User_id)

		user.Token = &token
		user.Refresh_token = &refreshToken

		_, insertErr := userCollection.InsertOne(ctx, user)

		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		user.Password = nil
		c.JSON(http.StatusOK, user)

	}
}

func HashPassword(password string) string {

	// to hash a password with cost factor 14
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}
	return string(pass)
}

func VerifyPassword(providedPassword string, userPassword string) (bool, string) {

	// to compare the provided password and the stored user password
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintln("password is incorrect")
		check = false
	}
	return check, msg
}

package handlers

import (
	"context"
	"net/http"
	"time"

	"bkawk/go-echo/api/models"
	"bkawk/go-echo/api/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// RegisterEndpoint handles user registration requests
func LoginPost(c echo.Context) error {
	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate input
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Find user by email or username
	var user models.User
	collection := db.Collection("users")
	err := collection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"email": u.Email},
			{"username": u.Email},
		},
	}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Check if password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Username, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

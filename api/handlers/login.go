package handlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"bkawk/go-echo/api/models"
	"bkawk/go-echo/api/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LoginStr struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

// RegisterEndpoint handles user registration requests
func LoginPost(c echo.Context) error {

	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate input
	u := new(LoginStr)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to bind request body"})
	}

	// Find user by email or username
	var user models.User
	collection := db.Collection("users")
	err := collection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"email": u.Email},
			{"username": u.Username},
		},
	}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Check if password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(os.Getenv("BCRYPT_PASSWORD")))
	if err != nil {

		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid brcypt credentials"})
	}

	// Generate JWT token
	jwtToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	// Generate the timestamp
	currentTime := time.Now().Unix()

	// Check if user already has a refresh token, if not update with one and last seen.
	var refreshToken string
	if user.RefreshToken == "" {
		// Generate refresh token
		resToken, err := utils.GenerateRefreshToken()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate refresh token"})
		}
		refreshToken = resToken

		// Update user with refresh token
		_, err = collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{
			"$set": bson.M{"refreshToken": refreshToken, "lastSeen": currentTime},
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate refresh token"})
		}
	} else {
		// If they do have a refresh token, update just the last seen date
		refreshToken = user.RefreshToken
		_, err = collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{
			"$set": bson.M{"lastSeen": currentTime},
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update last seen date"})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"jwtToken": jwtToken, "refreshToken": refreshToken})
}

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
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// RegisterEndpoint handles user registration requests
func LoginPost(c echo.Context) error {

	// Get logger from context
	logger := c.Get("logger").(*zap.Logger)
	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate input
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Generate the timestamp
	currentTime := time.Now().Unix()

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
		logger.Error("Failed to find user", zap.Error(err))
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Check if password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		logger.Error("Failed to compare password", zap.Error(err))
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	jwtToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		logger.Error("Failed to generate token", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	// Check if user already has a refresh token, if not update with one and last seen.
	var refreshToken string
	if user.RefreshToken == "" {
		// Generate refresh token
		refreshToken, err := utils.GenerateRefreshToken()
		if err != nil {
			logger.Error("Failed to generate refresh token", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate refresh token"})
		}

		// Update user with refresh token
		_, err = collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{
			"$set": bson.M{"refreshToken": refreshToken, "lastSeen": currentTime},
		})
		if err != nil {
			logger.Error("Failed to generate refresh token", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate refresh token"})
		}
	} else {
		// If they do have a refresh token, update just the last seen date
		refreshToken = user.RefreshToken
		_, err = collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{
			"$set": bson.M{"lastSeen": currentTime},
		})
		if err != nil {
			logger.Error("Failed to update last seen date", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update last seen date"})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"jwtToken": jwtToken, "refreshToken": refreshToken})
}

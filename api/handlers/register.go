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
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// RegisterEndpoint handles user registration requests
func RegisterPost(c echo.Context) error {

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

	// Validate strong password
	if err := utils.ValidatePassword(u.Password); err != nil {
		logger.Error("Password is not strong enough", zap.Error(err))
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Check if username and email are unique
	var existingUser models.User
	collection := db.Collection("users")
	err := collection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"username": u.Username},
			{"email": u.Email},
		},
	}).Decode(&existingUser)
	if err == nil {
		logger.Error("Username or Email already exists", zap.Error(err))
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username or Email already exists"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("BCRYPT_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash password", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}
	u.Password = string(hashedPassword)

	// Generate a unique user ID prefixed with "usr_"
	u.ID = "usr_" + utils.GenerateUUID()

	// Generate the timestamp
	u.CreatedAt = time.Now().Unix()

	// Save the user to MongoDB Atlas
	u.Password = string(hashedPassword)
	_, err = collection.InsertOne(ctx, u)
	if err != nil {
		logger.Error("Failed to save user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save user"})
	}

	// // Call the WelcomeEmail function
	// err = email.WelcomeEmail(u)
	// if err != nil {
	// 	logger.Error("We couldn't send a welcome email at this time, but your account has been successfully created", zap.Error(err))
	// 	return c.JSON(http.StatusOK, echo.Map{"error": "We couldn't send a welcome email at this time, but your account has been successfully created"})
	// }

	return c.JSON(http.StatusOK, echo.Map{"message": "Your account has been successfully created"})
}

package handlers

import (
	"bkawk/go-echo/api/customMiddleware"
	"bkawk/go-echo/api/emails"
	"bkawk/go-echo/api/utils"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserStr struct {
	ID               string `json:"_id,omitempty"`
	Username         string `json:"username" validate:"required,min=4,max=32"`
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,min=8,max=64"`
	Name             string `json:"name" validate:"required,min=1,max=64"`
	IsVerified       bool   `json:"isVerified"`
	VerificationCode string `json:"verificationCode,omitempty"`
	CreatedAt        int64  `json:"createdAt,omitempty"`
}

// RegisterEndpoint handles user registration requests
func RegisterPost(c echo.Context) error {

	// Validate input
	u := new(UserStr)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, customMiddleware.Response{Message: "Invalid request data"})
	}

	// Validate strong password
	if err := utils.CheckPasswordStrength(u.Password); err != nil {
		// Type assert the error to utils.PasswordError
		passwordError := err.(*utils.PasswordError)

		return c.JSON(http.StatusBadRequest, customMiddleware.Response{
			Message: "password not strong enough",
			Error: &customMiddleware.ErrorResponse{
				Errors: map[string]string{
					"password": passwordError.Password,
				},
			},
		})
	}

	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if username and email are unique
	var existingUser UserStr
	collection := db.Collection("users")
	err := collection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"username": u.Username},
			{"email": u.Email},
		},
	}).Decode(&existingUser)

	if err != nil && err != mongo.ErrNoDocuments {
		c.Logger().Errorf("Error querying database: %v", err)
		return c.JSON(http.StatusInternalServerError, customMiddleware.Response{Message: "An error occurred while processing your request"})
	}

	if err == nil {
		// Prepare the error response
		errorResponse := &customMiddleware.ErrorResponse{}

		if existingUser.Username == u.Username {
			errorResponse.Errors["username"] = "Username already exists"
		}
		if existingUser.Email == u.Email {
			errorResponse.Errors["email"] = "Email already exists"
		}

		return c.JSON(http.StatusBadRequest, customMiddleware.Response{Message: "Username or Email already exists", Error: errorResponse})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Logger().Errorf("Error hashing password: %v", err)
		return c.JSON(http.StatusInternalServerError, customMiddleware.Response{Message: "An error occurred while processing your request"})
	}
	u.Password = string(hashedPassword)

	// Generate a unique user ID prefixed with "usr_"
	uuid, err := utils.GenerateUUID()
	if err != nil {
		c.Logger().Errorf("Error generating user ID: %v", err)
		return c.JSON(http.StatusInternalServerError, customMiddleware.Response{Message: "An error occurred while processing your request"})
	}
	u.ID = "usr_" + uuid
	u.IsVerified = false

	// Generate a verification prefixed with "ver_"
	vCode, err := utils.GenerateUUID()
	if err != nil {
		c.Logger().Errorf("Error generating verification code: %v", err)
		return c.JSON(http.StatusInternalServerError, customMiddleware.Response{Message: "An error occurred while processing your request"})
	}
	u.VerificationCode = "ver_" + vCode

	// Generate the timestamp
	u.CreatedAt = time.Now().Unix()

	// Save the user to MongoDB Atlas
	_, err = collection.InsertOne(ctx, u)
	if err != nil {
		c.Logger().Errorf("Error saving user: %v", err)
		return c.JSON(http.StatusInternalServerError, customMiddleware.Response{Message: "An error occurred while processing your request"})
	}

	// Get the verification URL from the environment
	verifyUrl, exists := os.LookupEnv("VERIFY_URL")
	if !exists {
		c.Logger().Errorf("environment variable not set: VERIFY_URL")
		return c.JSON(http.StatusInternalServerError, customMiddleware.Response{Message: "An error occurred while processing your request"})
	}

	// Send welcome email
	emailError := emails.SendWelcomeEmail(u.Email, verifyUrl+"?verificationCode="+u.VerificationCode)
	if emailError != nil {
		c.Logger().Errorf("Error sending welcome email: %v", emailError)
		return c.JSON(http.StatusInternalServerError, customMiddleware.Response{Message: "An error occurred while processing your request"})
	}

	return c.JSON(http.StatusOK, customMiddleware.Response{Message: "Your account has been successfully created"})
}

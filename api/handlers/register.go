package handlers

import (
	"bkawk/go-echo/api/emails"
	"bkawk/go-echo/api/models"
	"bkawk/go-echo/api/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// RegisterEndpoint handles user registration requests
func RegisterPost(c echo.Context) error {

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
	if err := utils.CheckPasswordStrength(u.Password); err != nil {
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

		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username or Email already exists"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("BCRYPT_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}
	u.Password = string(hashedPassword)

	// Generate a unique user ID prefixed with "usr_"
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate user ID"})
	}
	u.ID = "usr_" + uuid
	u.IsVerified = false

	// Generate a verification prefixed with "ver_"
	vCode, err := utils.GenerateUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate user ID"})
	}
	u.VerificationCode = "ver_" + vCode

	// Generate the timestamp
	u.CreatedAt = time.Now().Unix()

	// Save the user to MongoDB Atlas
	u.Password = string(hashedPassword)
	_, err = collection.InsertOne(ctx, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save user"})
	}

	// Get the verification URL from the environment
	verifyUrl := os.Getenv("VERIFY_URL")
	if verifyUrl == "" {
		return fmt.Errorf("environment variable not set: VERIFY_URL")
	}

	// Send welcome email
	emailError := emails.SendWelcomeEmail(u.Email, verifyUrl+"?verificationCode="+u.VerificationCode)
	if emailError != nil {
		fmt.Println(emailError)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": emailError})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Your account has been successfully created"})
}

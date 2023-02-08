package handlers

import (
	"bkawk/go-echo/api/emails"
	"bkawk/go-echo/api/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User struct for storing user data in the database
type User struct {
	Email            string    `bson:"email"`
	VerificationCode string    `bson:"verificationCode"`
	CreatedAt        time.Time `bson:"createdAt"`
}

// Request struct for validating the email input
type Request struct {
	Email string `json:"email" validate:"required,email,max=100"`
}

// RegisterEndpoint handles user registration requests
func ResendVerificationPost(c echo.Context) error {
	// Bind request to the Request struct and validate the email
	request := new(Request)
	if err := c.Bind(request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	email := request.Email

	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if email exists in the users collection in MongoDB Atlas
	collection := db.Collection("users")
	var u User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusBadRequest, "Email not found")
		}
		return err
	}
	// Check if the CreatedAt timestamp is less than 10 minutes ago
	createdAt := u.CreatedAt
	if time.Since(createdAt) < 10*time.Minute {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Verification code already sent. Please wait %d minutes before trying again.", 10-int(time.Since(createdAt).Minutes())))
	}

	// Generate a verification prefixed with "ver_"
	vCode, err := utils.GenerateUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate user ID"})
	}
	u.VerificationCode = "ver_" + vCode

	// Generate the timestamp
	u.CreatedAt = time.Now()

	// Update user data in the collection
	_, err = collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"verificationCode": u.VerificationCode, "createdAt": u.CreatedAt}})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update user record"})
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

	return c.String(http.StatusOK, "Verification code sent")

}

package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"bkawk/go-echo/api/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// RegisterEndpoint handles user registration requests
func ResetPasswordPost(c echo.Context) error {
	var err error

	// Validate input
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to bind request body"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("BCRYPT_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to hash password")
	}

	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection("users")

	// Find the user document with the password reset token
	var user models.User
	fmt.Println(u.PasswordResetToken)
	err = collection.FindOne(ctx, bson.M{"passwordResetToken": u.PasswordResetToken}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.String(http.StatusBadRequest, "Invalid password reset token")
		}
		return c.String(http.StatusInternalServerError, "Error while finding user")
	}

	// Check if the password reset was requested within the past 24 hours
	forgotPassword := user.ForgotPassword
	delta := time.Now().Unix() - forgotPassword
	if delta > 24*60*60 {
		return c.String(http.StatusBadRequest, "Password reset token expired "+time.Duration(delta).String())
	}

	// Update the user document with the new password
	filter := bson.M{"passwordResetToken": u.PasswordResetToken}
	var update bson.M
	if time.Now().Unix()-forgotPassword > 5*60 {
		update = bson.M{"$set": bson.M{"password": string(hashedPassword)}, "$unset": bson.M{"passwordResetToken": ""}}
	} else {
		update = bson.M{"$set": bson.M{"password": string(hashedPassword)}}
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error while updating user")
	}

	return c.String(http.StatusOK, "Password successfully reset")
}

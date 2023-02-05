package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// VerifyEmailGet handles email verification GET requests
func VerifyEmailGet(c echo.Context) error {
	// Retrieve the verification code from the query parameter of the URL
	verificationCode := c.QueryParam("verificationCode")

	// Get the database connection from the context
	db := c.Get("db").(*mongo.Database)
	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the "users" collection from the database
	collection := db.Collection("users")
	// Define the filter to find the user with the given verification code
	filter := bson.M{"verificationCode": verificationCode}
	// Define the update to set the "isVerified" field to true
	update := bson.M{"$set": bson.M{"isVerified": true}}

	// Execute the update on the user in the database
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to verify the user"})
	}

	if res.ModifiedCount == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid verification code"})
	}

	// Return a JSON response with a success message
	return c.JSON(http.StatusOK, map[string]string{"message": "User marked as verified successfully!"})
}

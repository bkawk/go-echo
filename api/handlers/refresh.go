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
)

// Response is a struct to hold the response data for an HTTP request
type PostResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// RegisterEndpoint handles user registration requests
func RefreshPost(c echo.Context) error {
	// Get the refreshToken from the request
	refreshToken := c.FormValue("refreshToken")
	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the user with the matching refreshToken
	var user models.User
	err := db.Collection("users").FindOne(ctx, bson.M{"refreshToken": refreshToken}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusBadRequest, PostResponse{
				Message: "Invalid refresh token",
			})
		}
		return c.JSON(http.StatusInternalServerError, PostResponse{
			Message: "Failed to find user in database",
		})
	}

	// Generate a JWT with the user's ID added as a claim
	jwt, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, PostResponse{
			Message: "Failed to generate JWT",
		})
	}

	// Return the JWT to the user
	return c.JSON(http.StatusOK, PostResponse{
		Message: "Success",
		Data: map[string]string{
			"jwt": jwt,
		},
	})
}

type DeleteResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// RegisterEndpoint handles user registration requests
func RefreshDelete(c echo.Context) error {

	// Get the refreshToken from the request
	refreshToken := c.FormValue("refreshToken")

	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the user with the matching refreshToken and update it to an empty string
	_, err := db.Collection("users").UpdateOne(ctx, bson.M{"refreshToken": refreshToken}, bson.M{"$set": bson.M{"refreshToken": ""}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusBadRequest, DeleteResponse{
				Message: "Invalid refresh token",
			})
		}
		return c.JSON(http.StatusInternalServerError, DeleteResponse{
			Message: "Failed to update user in database",
		})
	}

	return c.JSON(http.StatusOK, DeleteResponse{
		Message: "Success",
	})
}

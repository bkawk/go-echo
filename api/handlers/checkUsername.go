package handlers

import (
	"bkawk/go-echo/api/models"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUsernameGet(c echo.Context) error {

	// Get the username from the query parameter
	username := c.QueryParam("username")

	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if username exists in the database
	var existingUser models.User
	collection := db.Collection("users")
	err := collection.FindOne(ctx, bson.M{
		"username": username,
	}).Decode(&existingUser)
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{"message": "Username not available", "isAvailable": false})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Username available", "isAvailable": true})
}

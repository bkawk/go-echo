package handlers

import (
	"bkawk/go-echo/api/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterEndpoint handles user registration requests
func ForgotPasswordPost(c echo.Context) error {

	// Get database connection from context
	db := c.Get("db").(*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate input
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to bind request body"})
	}

	var user models.User
	collection := db.Collection("users")
	filter := bson.M{"email": u.Email}

	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, "Email not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching user")
	}

	if time.Since(time.Unix(user.ForgotPassword, 0)) < (5 * time.Minute) {
		waitTime := 5*time.Minute - time.Since(time.Unix(user.ForgotPassword, 0))
		return echo.NewHTTPError(http.StatusTooManyRequests, fmt.Sprintf("Try again in %d minutes and %d seconds", int(waitTime.Minutes()), int(waitTime.Seconds())%60))
	}

	// send email logic
	fmt.Println("Sending forgot password email to", u.Email)

	user.ForgotPassword = time.Now().Unix()
	if _, err := collection.ReplaceOne(ctx, filter, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating user")
	}

	return c.JSON(http.StatusOK, "Forgot password email sent!")
}

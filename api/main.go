package main

import (
	"bkawk/go-echo/api/database"
	"bkawk/go-echo/api/handlers"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	client, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Disconnect(context.TODO())

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", client)
			return next(c)
		}
	})

	e.GET("/health", handlers.HealthGet)                            // Health Check
	e.POST("/register", handlers.RegisterPost)                      // Register a new user
	e.POST("/login", handlers.LoginPost)                            // Login with a username and password
	e.POST("/refresh", handlers.RefreshPost)                        // Refresh the access token using a refresh token
	e.DELETE("/refresh", handlers.RefreshDelete)                    // Invalidate the current refresh token, effectively logging the user out
	e.POST("/forgot-password", handlers.ForgotPasswordPost)         // Send a password reset email to the user
	e.POST("/reset-password", handlers.ResetPasswordPost)           // Reset the password of a user
	e.PUT("/profile", handlers.ProfilePut)                          // Update the profile information of a user
	e.GET("/profile/:username", handlers.ProfileGet)                // Retrieve the profile information of a specific user,
	e.GET("/verify-email", handlers.VerifyEmailGet)                 // Verify the email address of a user
	e.POST("/resend-verification", handlers.ResendVerificationPost) // Resend the verification email to a user
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

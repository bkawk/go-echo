package main

import (
	"bkawk/go-echo/api/database"
	"bkawk/go-echo/api/handlers"
	"context"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/juju/ratelimit"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Initialize Echo
	e := echo.New()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Connect to MongoDB Atlas
	client, err := database.Connect()
	if err != nil {
		return
	}
	defer client.Disconnect(context.TODO())

	// Limit the number of requests to 1 requests per second with a burst of 20 requests
	limiter := ratelimit.NewBucketWithQuantum(1, 20, 1)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if limiter.TakeAvailable(1) == 0 {
				msg := "Rate limit exceeded"
				return c.String(http.StatusTooManyRequests, msg)
			}
			return next(c)
		}
	})

	// Inject MongoDB client into Echo using middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", client.Database(os.Getenv("MONGO_DB")))
			return next(c)
		}
	})

	// Configure CORS middleware
	config := middleware.CORSConfig{
		// AllowOrigins: []string{"http://www.mydomain.com"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}
	e.Use(middleware.CORSWithConfig(config))

	// Limit Body Size
	e.Use(middleware.BodyLimit("100K"))

	// Add middleware to be more secure with xss etc
	e.Use(middleware.Secure())

	// Routes
	e.GET("/health", handlers.HealthGet)                            // Health Check
	e.POST("/register", handlers.RegisterPost)                      // Register a new user
	e.GET("/check-username", handlers.CheckUsernameGet)             // Check if a username is available
	e.POST("/login", handlers.LoginPost)                            // Login with a username and password
	e.POST("/refresh", handlers.RefreshPost)                        // Refresh the access token using a refresh token
	e.DELETE("/refresh", handlers.RefreshDelete)                    // Invalidate the current refresh token, effectively logging the user out
	e.POST("/forgot-password", handlers.ForgotPasswordPost)         // Send a password reset email to the user
	e.POST("/reset-password", handlers.ResetPasswordPost)           // Reset the password of a user
	e.PUT("/profile", handlers.ProfilePut)                          // Update the users profile information
	e.GET("/profile/:username", handlers.ProfileGet)                // Retrieve the public profile information of a specific user,
	e.GET("/verify-email", handlers.VerifyEmailGet)                 // Verify the email address of a user
	e.POST("/resend-verification", handlers.ResendVerificationPost) // Resend the verification email to a user

	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

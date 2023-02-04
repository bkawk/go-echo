package main

import (
	"bkawk/go-echo/api/database"
	"bkawk/go-echo/api/handlers"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/juju/ratelimit"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Connect to MongoDB Atlas
	client, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Disconnect(context.TODO())

	// Initialize Echo
	e := echo.New()

	// Inject MongoDB client into Echo using middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", client.Database(os.Getenv("MONGO_DB")))
			return next(c)
		}
	})

	// Limit the number of requests to 1 requests per second with a burst of 20 requests
	limiter := ratelimit.NewBucketWithQuantum(1, 20, 1)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if limiter.TakeAvailable(1) == 0 {
				return c.String(http.StatusTooManyRequests, "Rate limit exceeded")
			}
			return next(c)
		}
	})

	// Zap logger
	// logger.Debug("debug log")
	// logger.Info("info log")
	// logger.Warn("warn log")
	// logger.Error("error log")
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			latency := stop.Sub(start)

			fields := []zap.Field{
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.Int("status", res.Status),
				zap.Duration("latency", latency),
				zap.String("remote_ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
			}

			if res.Status >= http.StatusBadRequest {
				logger.Warn("Request failed", fields...)
			} else {
				logger.Info("Request succeeded", fields...)
			}

			return nil
		}
	})

	// Routes
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

	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

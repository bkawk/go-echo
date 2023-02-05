package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// BearerTokenMiddleware is a middleware function to check for a bearer token
func Route(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the authorization header from the request
		authHeader := c.Request().Header.Get("Authorization")

		// Check if the header is empty
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}

		// Split the header value into an array of two strings
		// where the first string is the token type and the second is the token itself
		split := strings.Split(authHeader, " ")
		if len(split) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is invalid")
		}

		// Check if the token type is "Bearer"
		if split[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is invalid")
		}

		// Get the token value
		token := split[1]

		// Validate the token (in this example, a simple string comparison is used)
		if token != "yoursecretbearertoken" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token is invalid")
		}

		// Call the next handler
		return next(c)
	}
}

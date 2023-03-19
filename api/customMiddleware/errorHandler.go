package customMiddleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Errors map[string]string `json:"errors,omitempty"`
}

type Response struct {
	Message string         `json:"message"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			c.Logger().Error(err)

			// Your custom error response
			errorResponse := &ErrorResponse{
				Errors: map[string]string{
					"error": err.Error(),
				},
			}

			response := &Response{
				Message: "An error occurred",
				Error:   errorResponse,
			}

			return c.JSON(http.StatusInternalServerError, response)
		}
		return nil
	}
}

package handlers

import (
	"net/http"

	"bkawk/go-echo/api/models"

	"github.com/labstack/echo/v4"
)

// RegisterEndpoint handles user registration requests
func ProfilePut(c echo.Context) error {
	// bind the incoming request body to a User struct
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// validate user input
	if u.Username == "" || u.Password == "" || u.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	// add the new user to the database
	// (this is a dummy implementation and would be replaced in a real application)
	// ...

	// return a success response
	return c.JSON(http.StatusOK, map[string]string{
		"message": "user registered successfully",
	})
}

// RegisterEndpoint handles user registration requests
func ProfileGet(c echo.Context) error {
	// bind the incoming request body to a User struct
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// validate user input
	if u.Username == "" || u.Password == "" || u.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	// add the new user to the database
	// (this is a dummy implementation and would be replaced in a real application)
	// ...

	// return a success response
	return c.JSON(http.StatusOK, map[string]string{
		"message": "user registered successfully",
	})
}

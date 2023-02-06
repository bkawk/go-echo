package customMiddleware

import (
	"bkawk/go-echo/api/utils"

	"github.com/labstack/echo/v4"
)

func Fingerprint(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		userAgent := c.Request().UserAgent()

		// Get device type and screen resolution
		deviceType := c.Request().Header.Get("Device-Type")
		screenResolution := c.Request().Header.Get("Screen-Resolution")

		// Get browser and operating system information
		browser := c.Request().Header.Get("Browser")
		os := c.Request().Header.Get("Operating-System")

		// Generate hashed fingerprint string
		fingerprint := userAgent + deviceType + screenResolution + browser + os
		hashedFingerprint, err := utils.Hash([]byte(fingerprint))
		if err != nil {
			return echo.NewHTTPError(500, "Error generating hashed fingerprint")
		}

		// Add the fingerprint to the context so that it can be used later in the request
		c.Set("fingerprint", hashedFingerprint)

		return next(c)

	}
}

// func MyHandler(c echo.Context) error {
// 	fingerprint := c.Get("fingerprint").(string)
// 	// Use the fingerprint in your handler logic
// 	...
// 	return c.JSON(http.StatusOK, ...)
// }

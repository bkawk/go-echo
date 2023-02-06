package customMiddleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func Fingerprint(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userAgent := c.Request().UserAgent()
		deviceType := getDeviceType(userAgent)
		screenResolution := getScreenResolution(c.Request().Header.Get("User-Agent"))

		fingerprint := deviceType + "-" + screenResolution

		// Add the fingerprint to the context so that it can be used later in the request
		c.Set("fingerprint", fingerprint)

		return next(c)
	}
}

func getDeviceType(userAgent string) string {
	if strings.Contains(userAgent, "Mobile") {
		return "mobile"
	}
	return "desktop"
}

func getScreenResolution(userAgent string) string {
	// Example implementation, actual implementation would depend on the specific requirements
	return "1920x1080"
}

// func MyHandler(c echo.Context) error {
// 	fingerprint := c.Get("fingerprint").(string)
// 	// Use the fingerprint in your handler logic
// 	...
// 	return c.JSON(http.StatusOK, ...)
// }

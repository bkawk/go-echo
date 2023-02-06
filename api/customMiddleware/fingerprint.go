package customMiddleware

import (
	"bkawk/go-echo/api/utils"

	"github.com/labstack/echo/v4"
)

func Fingerprint(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Collect information from request headers
		userAgent := c.Request().UserAgent()
		acceptLanguage := c.Request().Header.Get("Accept-Language")
		acceptEncoding := c.Request().Header.Get("Accept-Encoding")

		// Get the combined cookie
		fpCookie, err := c.Cookie("fp")
		if err != nil {
			return next(c)
		}

		// Generate hashed fingerprint string
		fingerprint := userAgent + acceptLanguage + acceptEncoding + fpCookie.Value
		hashedFingerprint, err := utils.Hash([]byte(fingerprint))
		if err != nil {
			// Error generating hashed fingerprint, fail silently
			return next(c)
		}

		// Add the fingerprint to the context so that it can be used later in the request
		c.Set("fingerprint", hashedFingerprint)

		return next(c)

	}
}

// hashedFingerprint := c.Get("fingerprint").(string)

// function gatherInfoAndSetCookie() {
//   // Gather information about the user's device
//   var screenWidth = window.screen.width;
//   var screenHeight = window.screen.height;
//   var installedFonts = window.fonts.length;
//   var timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;

//   // Combine the gathered information into a single string
//   var gatheredInfo = screenWidth + screenHeight + installedFonts + timezone;

//   // Set a cookie with the gathered information
//   document.cookie = "fp=" + gatheredInfo + "; expires=Fri, 31 Dec 9999 23:59:59 GMT; path=/;";
// }

// var w=window.screen,a=w.width,b=w.height,c=window.fonts.length,d=Intl.DateTimeFormat().resolvedOptions().timeZone,e=a+b+c+d;document.cookie="fp="+e+"; expires=Fri, 31 Dec 9999 23:59:59 GMT; path=/;";

// var _0x88d4=["\x73\x63\x72\x65\x65\x6E","\x77\x69\x64\x74\x68","\x68\x65\x69\x67\x68\x74","\x6C\x65\x6E\x67\x74\x68","\x66\x6F\x6E\x74\x73","\x74\x69\x6D\x65\x5A\x6F\x6E\x65","\x72\x65\x73\x6F\x6C\x76\x65\x64\x4F\x70\x74\x69\x6F\x6E\x73","\x63\x6F\x6F\x6B\x69\x65","\x66\x70\x3D","\x3B\x20\x65\x78\x70\x69\x72\x65\x73\x3D\x46\x72\x69\x2C\x20\x33\x31\x20\x44\x65\x63\x20\x39\x39\x39\x39\x20\x32\x33\x3A\x35\x39\x3A\x35\x39\x20\x47\x4D\x54\x3B\x20\x70\x61\x74\x68\x3D\x2F\x3B"];var w=window[_0x88d4[0]],a=w[_0x88d4[1]],b=w[_0x88d4[2]],c=window[_0x88d4[4]][_0x88d4[3]],d=Intl.DateTimeFormat()[_0x88d4[6]]()[_0x88d4[5]],e=a+ b+ c+ d;document[_0x88d4[7]]= _0x88d4[8]+ e+ _0x88d4[9]

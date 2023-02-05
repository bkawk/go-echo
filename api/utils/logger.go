package utils

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

var (
	log *zap.Logger
)

func Init() {
	var err error
	log, err = zap.NewProduction()
	if err != nil {
		// handle error
	}
}

func LoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			uri := c.Request().RequestURI
			status := c.Response().Status

			logger.Info("request",
				zap.String("URI", uri),
				zap.Int("status", status),
			)

			return next(c)
		}
	}
}

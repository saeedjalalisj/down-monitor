package web

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// ServerHeader middleware adds a `Server` header to the response.
func Logger(log *zap.SugaredLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Infow("response started", "IP", c.RealIP(), "Method", c.Request().Method, "path", c.Request().URL.Path, "req-id", c.Request().Header.Get("X-Request-ID"))
			return next(c)
		}
	}
}

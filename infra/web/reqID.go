package web

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func ReqID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ID := uuid.New().String()
		c.Request().Header.Add("X-Request-ID", ID)
		return next(c)
	}
}


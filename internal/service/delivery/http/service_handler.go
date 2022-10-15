package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/saeedjalalisj/down-monitor/internal/domain"
)

type ResponseError struct {
	Message string `json:"message"`
}

type serviceHandler struct {
	SUsecase domain.ServiceUsecase
}

func NewServiceHandler(e *echo.Echo, su domain.ServiceUsecase) {
	handler := serviceHandler{
		SUsecase: su,
	}

	e.POST("/service", handler.Create)
}

func (h *serviceHandler) Create(c echo.Context) (err error) {
	var service domain.CreateServiceDto

	err = c.Bind(&service)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// var ok bool
	// if ok, err = isRequestValid(&service); !ok {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	ctx := c.Request().Context()
	_, err = h.SUsecase.Create(ctx, &service)
	if err != nil {
		return c.JSON(404, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, service)
}

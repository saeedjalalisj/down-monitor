package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/saeedjalalisj/down-monitor/infra/web"
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
		web.AddResponseToContext(c, web.RespObject{
			Message: "Error",
			Data:    err.Error(),
			Code:    http.StatusUnprocessableEntity,
		})
		return nil
	}
	if err = c.Validate(service); err != nil {
		web.AddResponseToContext(c, web.RespObject{
			Message: "Error",
			Data:    err.Error(),
			Code:    http.StatusBadRequest,
		})
		return nil
	}
	ctx := c.Request().Context()
	id, err := h.SUsecase.Create(ctx, &service)
	if err != nil {
		web.AddResponseToContext(c, web.RespObject{
			Message: "Error",
			Data:    err.Error(),
			Code:    http.StatusBadRequest,
		})
		return nil
	}

	web.AddResponseToContext(c, web.RespObject{
		Message: "create",
		Data:    id,
		Code:    http.StatusOK,
	})

	return nil
}

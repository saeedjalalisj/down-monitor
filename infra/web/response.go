package web

import "github.com/labstack/echo"

type RespObject struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

const (
	echoCtxRespKey = "echo_ctx_resp_key"
)

func AddResponseToContext(c echo.Context, response RespObject) {
	c.Set(echoCtxRespKey, response)
}

func GetResponseFromContext(c echo.Context) RespObject {
	return c.Get(echoCtxRespKey).(RespObject)
}

func ResponseMiddle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		next(c)
		resp := GetResponseFromContext(c)
		return c.JSON(resp.Code, resp)
	}
}

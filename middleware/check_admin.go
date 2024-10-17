package middleware

import (
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"net/http"

	"github.com/labstack/echo"
)

func ISAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// handler login
			req := req.ReqSignUp{}
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			if req.Email != "admin@gmail.com" {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "bạn không có quyền gọi api này",
					Data:       nil,
				})
			}
			return next(c)
		}
	}
}

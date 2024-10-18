package middleware

import (
	"backend-github-trending/model"
	"backend-github-trending/security"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// JWTMiddleware tạo middleware để xác thực JWT
func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{}, // Sử dụng secret key từ security package
		SigningKey: security.SECRECT_KEY,     // Sử dụng struct JwtCustomClaims từ model
	}
	return middleware.JWTWithConfig(config)
}

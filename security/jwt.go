package security

import (
	"backend-github-trending/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRECT_KEY = "vhoang10293"

// func JWTMiddleware
func Gentoken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRECT_KEY))
	if err != nil {
		return "", err
	}

	return result, nil
}

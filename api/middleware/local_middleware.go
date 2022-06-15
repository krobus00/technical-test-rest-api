package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/labstack/echo/v4"
)

func VerifyToken(tokenSecret, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeJWTTokenMiddleware(tokenSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			tokenHeader := eCtx.Request().Header.Get("Authorization")
			tokenHeader = strings.Replace(tokenHeader, "Bearer ", "", -1)
			if tokenHeader == "" {
				return model.NewHttpCustomError(http.StatusUnauthorized, errors.New("Invalid Token"))
			}

			token, err := VerifyToken(tokenSecret, tokenHeader)
			if err != nil {
				return model.NewHttpCustomError(http.StatusUnauthorized, errors.New("Invalid Token"))
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				userID, ok := claims["userID"].(string)
				if !ok {
					return model.NewHttpCustomError(http.StatusUnauthorized, errors.New("Invalid Token"))
				}
				eCtx.Set("userID", userID)
			} else {
				return model.NewHttpCustomError(http.StatusUnauthorized, errors.New("Invalid Token"))
			}
			return next(eCtx)
		}
	}
}

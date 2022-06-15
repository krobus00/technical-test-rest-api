package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/krobus00/technical-test-rest-api/model"
)

func CreateToken(userID string, exp time.Duration, secret string) (*model.Token, error) {
	tokenExp := time.Now().Add(exp).Unix()
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userID"] = userID
	atClaims["exp"] = tokenExp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &model.Token{
		Token: token,
		Exp:   tokenExp,
	}, nil
}

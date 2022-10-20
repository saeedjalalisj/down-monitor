package jwtService

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func GenerateJWT(serviceName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(3600 * time.Minute)
	claims["service"] = serviceName
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "Signing Error", err
	}
	return tokenString, nil
}

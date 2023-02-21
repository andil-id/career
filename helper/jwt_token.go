package helper

import (
	"career/config"
	"career/model/web"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenereateJwtToken(id string, name string) (string, error) {
	claims := web.Claims{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Career",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt_secret := config.JwtSecreet()
	signedToken, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

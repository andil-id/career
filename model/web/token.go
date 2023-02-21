package web

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

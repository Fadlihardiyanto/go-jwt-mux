package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("akshdkjahsdi2138y1y8810")

type JWTClaims struct {
	Username string
	jwt.RegisteredClaims
}

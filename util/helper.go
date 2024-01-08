package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("secret") // Ganti "your-secret-key" dengan kunci rahasia yang aman

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString(SecretKey)
}

func Parsejwt(cookie string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, jwt.ErrInvalidClaimsType // Sesuaikan dengan error handling yang sesuai
	}

	return claims, nil
}

package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(getSecret())

func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "default_key"
	}
	return secret
}

func GenerateToken(userID uint, name string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"name":    name,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (uint, string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("invalid claims")
	}

	idFloat, okID := claims["user_id"].(float64)
	name, okName := claims["name"].(string)

	if !okID || !okName {
		return 0, "", errors.New("missing claims")
	}

	return uint(idFloat), name, nil
}

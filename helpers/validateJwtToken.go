package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func validateJwtToken(tokenStr string) bool {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validating the Signing Method ---
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		return time.Now().After(expirationTime)
	}

	return false
}

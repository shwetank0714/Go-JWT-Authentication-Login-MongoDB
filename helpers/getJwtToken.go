package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("secret_key")

func GetJwtToken(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours)
		"iat":     time.Now().Unix(),                     // Token issuance time
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", errors.New("failed to sign the JWT token")
	}

	return signedToken, nil
}

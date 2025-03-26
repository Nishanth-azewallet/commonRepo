// common/jwt.go
package common

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key") // Secret key used for signing and verifying JWT tokens

// GenerateJWT generates a JWT token with the userID and expiration time
func GenerateJWT(userID string, expirationTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(expirationTime).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseJWT validates and parses the JWT token, returning the user ID if valid
func ParseJWT(tokenStr string) (string, error) {
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Ensure token's signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("unable to extract claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("missing user ID in token")
	}

	return userID, nil
}

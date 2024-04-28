package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/MorrisMorrison/retfig/infrastructure/config"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string, issuer string, duration time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		Issuer:    issuer,
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := getSecret()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	secretKey := getSecret()

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func getSecret() string {
	return config.GetEnv("RETFIG_JWT_SECRET", "testsecret")
}

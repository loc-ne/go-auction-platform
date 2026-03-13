package pkg

import (
	"context"
	"fmt"
	"time"
	"crypto/rand"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	FullName string `json:"full_name"`
	jwt.RegisteredClaims
}

func GenerateTokens(ctx context.Context, userID string, email string, fullName string, jwtSecret string) (string, string, error) {
	secret := []byte(jwtSecret)
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		FullName: fullName,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID, 
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		return "", "", err
	}

	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", err
	}
	refreshToken := hex.EncodeToString(b)

	return accessToken, refreshToken, nil
}

func ValidateToken(tokenString string, jwtSecret string) (*JWTClaims, error) {
	secret := []byte(jwtSecret)
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, fmt.Errorf("invalid token")
}
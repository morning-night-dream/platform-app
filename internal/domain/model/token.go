package model

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	IDTokenKey      = "id-token"
	SessionTokenKey = "session-token"
	IDKey           = "id"
	SignKey         = "secret"
	DefaultExpires  = 86400 // 1 day
)

type ExpiresIn time.Duration

func GenerateToken(id string, sign string) (string, error) {
	claims := jwt.MapClaims{
		IDKey: id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	strToken, err := token.SignedString([]byte(sign))
	if err != nil {
		return "", err
	}

	return strToken, nil
}

func ValidToken(token string, sign string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(sign), nil
	})
	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func GetID(token string, sign string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(sign), nil
	})
	if err != nil {
		return "", err
	}

	if !parsedToken.Valid {
		return "", fmt.Errorf("invalid token")
	}

	parsedClaims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("couldn't parse claims")
	}

	uid, ok := parsedClaims[IDKey].(string)
	if !ok {
		return "", fmt.Errorf("couldn't parse claims")
	}

	return uid, nil
}

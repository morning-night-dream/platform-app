package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const (
	IDTokenKey = "uid"
	SignKey    = "secret"
)

func main() {
	uid := uuid.NewString()

	claims := jwt.MapClaims{
		IDTokenKey: uid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SignKey))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", tokenString)

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SignKey), nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", parsedToken)

	parsedClaims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		panic("couldn't parse claims")
	}

	if !parsedToken.Valid {
		panic("invalid token")
	}

	fmt.Printf("uid: %s", parsedClaims[IDTokenKey])
}

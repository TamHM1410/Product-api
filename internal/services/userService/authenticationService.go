package userService

import (
	"fmt"
	"go-backend-api/internal/services/redisService"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("mySecretKey")

func verifiedToken(tokenString string) any {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println("Token invalid:", err)
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token payload:", claims)
		fmt.Println("UserID:", claims["sub"])
		fmt.Println("Role:", claims["role"])
	} else {
		fmt.Println("Token invalid")
	}

	return token
}

func AuthenticateUser(token string) bool {

	if token == "" {
		return false
	}

	cachedToken := redisService.GetFromUserTokenKey(token)

	if cachedToken == nil {
		return false
	}

	verify := verifiedToken(token)

	if verify == nil {
		return false
	}

	return true
}

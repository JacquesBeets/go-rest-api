package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "supersectret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return nil, errors.New("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return claims, nil
}

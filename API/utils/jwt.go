package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretstring"

func GenerateToken(email, userId, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"userId":   userId,
		"userType": userType,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) (bool, string, string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("error validating JWT")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, "", "", errors.New("could not parse token")
	}

	isValid := parsedToken.Valid

	if !isValid {
		return false, "", "", errors.New("could not parse token")
	}

	//Extracting JWT content

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return false, "", "", errors.New("could not parse token")
	}

	userType := claims["userType"].(string)
	userId := claims["userId"].(string)

	return true, userType, userId, nil
}

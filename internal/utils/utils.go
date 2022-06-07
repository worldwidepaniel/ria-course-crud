package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetPasswordHash(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 0)
	return string(hash), err
}

func GenerateJWT(userEmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":     userEmail,
		"valid_to": time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func jwtParser(jwtString string) *jwt.Token {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil
	}
	return token
}

func IsJWTValid(jwtString string) bool {
	token := jwtParser(jwtString)
	if token == nil {
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	valid_to := fmt.Sprintf("%v", claims["valid_to"])
	valid_parsed, err := strconv.ParseFloat(valid_to, 64)
	if err != nil {
		return false
	}
	if ok && token.Valid && time.Unix(int64(valid_parsed), 0).After(time.Now()) {
		return true
	}
	return false

}

func UserEmailFromJWT(jwtString string) string {
	token := jwtParser(jwtString)
	if isValid := IsJWTValid(jwtString); isValid {
		claims, _ := token.Claims.(jwt.MapClaims)
		return fmt.Sprintf("%v", claims["user"])
	}
	return ""
}

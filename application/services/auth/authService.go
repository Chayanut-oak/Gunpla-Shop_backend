package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	secretKey = []byte(os.Getenv("SECRET_KEY"))
}

type AuthService struct{}

func (s *AuthService) GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "your-issuer",
	})

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) ValidateToken(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Print(token)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, ok := claims["sub"].(string)
		if !ok {
			return "", "", errors.New("user ID not found in token claims")
		}
		iss, ok := claims["iss"].(string)
		if !ok {
			return "", "", errors.New("iss not found in token claims")
		}
		fmt.Print(email, iss, nil)
		return email, iss, nil

	}

	return "", "", errors.New("invalid token claims")
}
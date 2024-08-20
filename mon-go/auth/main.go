package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))


func GenerateToken(email string, username string) (tokenString string, err error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err = token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type Claims struct {
	Username string `json:"username"`
	Email string `json:"email"`
}

func ValidateToken(tokenString string) (claims Claims,err error) {
	token, err := jwt.Parse(tokenString ,func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return claims,err
	}
	if !token.Valid {
		return claims,fmt.Errorf("invalid token")
	}

	res := token.Claims.(jwt.MapClaims);

	claims = Claims{
		Username: res["username"].(string),
		Email: res["email"].(string),
	}

	return claims, nil
}
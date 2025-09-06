package utils

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ValidateToken(r *http.Request) (interface{}, error) {
	token, err := r.Cookie("token")

	if err != nil {
		return nil, errors.New("cannot parse token")
	}

	parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("cannot parse token")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, errors.New("cannot parse token")
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return nil, errors.New("token timed out")
			}
		} else {
			return nil, errors.New("unknown error")
		}

		userEmail := claims["email"]
		userId := claims["id"]

		user := map[string]interface{}{
			"id":    userId,
			"email": userEmail,
		}

		return user, nil
	}
	return nil, errors.New("unknown error")
}

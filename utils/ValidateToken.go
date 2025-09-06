package utils

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserReturn struct {
	Email string `json:"email,omitempty" bson:"email"`
	Id    string `json:"id,omitempty" bson:"id"`
}

func ValidateToken(r *http.Request) (UserReturn, error) {
	token, err := r.Cookie("token")

	var user UserReturn

	if err != nil {
		return user, errors.New("cannot parse token")
	}

	parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("cannot parse token")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return user, errors.New("cannot parse token")
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return user, errors.New("token timed out")
			}
		} else {
			return user, errors.New("unknown error")
		}

		user.Email = claims["email"].(string)
		user.Id = claims["id"].(string)

		return user, nil
	}

	return user, errors.New("unknown error")
}

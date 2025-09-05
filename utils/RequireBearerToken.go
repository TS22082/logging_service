package utils

import (
	"errors"
	"net/http"
)

func RequireBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}
	return authHeader, nil
}

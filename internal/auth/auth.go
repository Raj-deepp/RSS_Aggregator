package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeader = errors.New("no authorization header included")

// Extracts API key from the Headers of HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed Authorization Header")
	}

	return splitAuth[1], nil
}

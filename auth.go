package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/vihaan404/aggreg/internal/database"
	"log/slog"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")
var ErrMalFormed = errors.New("malformed authorization header")

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (c *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := GetAPIKey(r.Header)
		if err != nil {
			fmt.Println(err)
		}
		i, err := c.db.GetUserApiKey(context.Background(), apikey)
		if err != nil {
			slog.Error("authentication failedd")
		}
		handler(w, r, i)

	}
}

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", ErrMalFormed
	}

	return splitAuth[1], nil
}

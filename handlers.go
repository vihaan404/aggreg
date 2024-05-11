package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/vihaan404/aggreg/internal/database"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type FeedBody struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *apiConfig) getFeed(w http.ResponseWriter, r *http.Request) {
	feeds, err := c.db.GetAllFeeds(context.Background())
	if err != nil {
		log.Println(err)
	}
	for _, feed := range feeds {
		respondWithJSON(w, http.StatusOK, feed)
	}
}
func (c *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedBody := FeedBody{}
	err := json.NewDecoder(r.Body).Decode(&feedBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	params := database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedBody.Name,
		Url:       feedBody.URL,
		UserID:    user.ID,
	}
	feed, err := c.db.CreateFeeds(context.Background(), params)
	if err != nil {
		log.Println("error creating feed:", err)
	}
	respondWithJSON(w, http.StatusCreated, feed)
}

func (c *apiConfig) readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		slog.Error("Error writing response")
	}
}
func (c *apiConfig) errHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(fmt.Sprint("Internal server error")))
	if err != nil {
		slog.Error("Error writing response")
	}

}

func (c *apiConfig) getUserApiHandler(w http.ResponseWriter, r *http.Request) {
	apikey, err := GetAPIKey(r.Header)
	if err != nil {
		fmt.Println(err)
	}
	i, err := c.db.GetUserApiKey(context.Background(), apikey)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	respondWithJSON(w, http.StatusOK, i)

}
func (c *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := RequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      requestBody.Name,
	}
	i, err := c.db.CreateUser(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	respondWithJSON(w, http.StatusOK, i)
}

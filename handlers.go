package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
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

func (c *apiConfig) getFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := c.db.GetFeedFollow(r.Context(), user.ID)

	if err != nil {
		log.Println(err)
	}
	for _, feedFollow := range feedFollows {
		respondWithJSON(w, http.StatusOK, feedFollow)
	}
}
func (c *apiConfig) getFeed(w http.ResponseWriter, r *http.Request) {

	feeds, err := c.db.GetAllFeeds(r.Context())
	if err != nil {
		log.Println(err)
	}
	for _, feed := range feeds {
		respondWithJSON(w, http.StatusOK, feed)
	}
}

type FeedFollowBody struct {
	Feed_id uuid.UUID `json:"feed_id"`
}

func (c *apiConfig) createFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowBody := FeedFollowBody{}
	err := json.NewDecoder(r.Body).Decode(&feedFollowBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feedFollowBody.Feed_id,
		UserID:    user.ID,
	}
	feed_follow, err := c.db.CreateFeedFollow(r.Context(), params)
	if err != nil {
		log.Println("error creating feed_follow:", err)
	}
	respondWithJSON(w, http.StatusCreated, feed_follow)

}
func (c *apiConfig) deleteFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	id := chi.URLParamFromCtx(r.Context(), "id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = c.db.DeleteFeedFollow(r.Context(), uuidID)
	if err != nil {
		log.Fatal("error deleting feed_follow:", err)
	}

	respondWithJSON(w, http.StatusOK, nil)

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

	feed, err := c.db.CreateFeeds(r.Context(), params)

	if err != nil {
		log.Println("error creating feed:", err)
	}
	paramsFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}
	_, err = c.db.CreateFeedFollow(r.Context(), paramsFollow)
	if err != nil {
		log.Println("error creating feed follow ", err)
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
	i, err := c.db.GetUserApiKey(r.Context(), apikey)
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
	i, err := c.db.CreateUser(r.Context(), params)
	if err != nil {
		log.Fatal(err)
	}

	respondWithJSON(w, http.StatusOK, i)
}

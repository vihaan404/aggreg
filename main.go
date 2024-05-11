package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vihaan404/aggreg/internal/database"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	godotenv.Load(".env")
	Port := os.Getenv("PORT")
	dbURL := os.Getenv("POSTGRES")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error connecting to the database", err)
	}
	dbQuries := database.New(db)
	api := &apiConfig{db: dbQuries}
	router := chi.NewRouter()
	router.Get("/v1/err", api.errHandler)
	router.Get("/v1/readiness", api.readinessHandler)
	router.Post("/v1/create", api.createUserHandler)
	router.Get("/v1/user", api.getUserApiHandler)
	router.Post("/v1/feed", api.middlewareAuth(api.createFeedHandler))
	router.Get("/v1/feeds", api.getFeed)

	srv := &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}
	fmt.Println("Listening on port " + Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	payloadBytes, _ := json.Marshal(payload)
	_, err := w.Write(payloadBytes)
	if err != nil {
		slog.Error("Error writing response")
	}

}

type RequestBody struct {
	Name string `json:"name"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write([]byte(msg))
	if err != nil {
		slog.Error("Error writing response")
	}
}

type status struct {
	status string
}

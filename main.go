package main

import (
	"log"
	"net/http"

	"github.com/dawidcyron/shortener/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler)
	database.NewRedisClient()
	router.Post("/shorten", ShortenURL)
	router.Get("/{id}", GetFullURL)
	log.Fatal(http.ListenAndServe(":8080", router))
}

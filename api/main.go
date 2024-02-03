package main

import (
	"challenge/api/controllers"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token, Authorization"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"message":"Welcome to the Enron-Email Index OpenObserve API"}`))
		})

		r.Get("/search/{word}", func(w http.ResponseWriter, r *http.Request) {
			searchTerm := chi.URLParam(r, "word")
			if searchTerm == "" {
				http.Error(w, `{"error":"Query parameter 'word' is required"}`, http.StatusBadRequest)
				return
			}
			result := controllers.SearchDocument(searchTerm)
			json, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(json)
		})
	})

	log.Println("Server is up and running on port 3001")
	http.ListenAndServe(":3001", r)
}

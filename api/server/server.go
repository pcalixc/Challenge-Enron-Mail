package server

import (
	"challenge/api/controllers"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type Server struct {
	ChiRouter *chi.Mux
}

func New() Server {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(cors.Handler(cors.Options{

		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	server := Server{
		ChiRouter: router,
	}

	server.setupRoutesb() // Configurar las rutas

	return server
}

func (s *Server) Run() error {
	port := os.Getenv("API_PORT")

	fmt.Println("running on:" + port)
	return http.ListenAndServe(":"+port, s.ChiRouter)
}

func (s *Server) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"Enron-Email Index ZincSearch API"}`))
}

func (s *Server) handleEmails(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("term")
	page := r.URL.Query().Get("page")
	max := r.URL.Query().Get("max")

	if searchTerm != "" {
		term, err := controllers.SearchMails(searchTerm, page, max)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		jsonData, err := json.Marshal(term)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	} else {
		result := controllers.GetAllEmails(page, max)
		jsonData, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
	}
}

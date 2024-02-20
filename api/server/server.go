package server

import (
	"challenge/api/config"
	"challenge/api/controllers"
	"challenge/api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
		MaxAge:           300,
	}))

	server := Server{
		ChiRouter: router,
	}

	server.setupRoutes()

	return server
}

func (s *Server) Run() error {
	port := config.GetEnvVar("API_PORT")

	fmt.Println("running on:" + port)
	return http.ListenAndServe(":"+port, s.ChiRouter)
}

func (s *Server) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"Welcome to Enron-Email Index ZincSearch API 💫"}`))
}

func (s *Server) handleEmails(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("term")
	page := r.URL.Query().Get("page")
	max := r.URL.Query().Get("max")

	intMax, err := strconv.Atoi(max)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to convert max_results to integer: %v", err), http.StatusBadRequest)
		return
	}

	intPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to convert page to integer: %v", err), http.StatusBadRequest)
		return

	}

	var result models.HitsResponse
	var jsonData []byte

	if searchTerm != "" {
		result, err = controllers.SearchMails(searchTerm, intPage, intMax)
	} else {
		result, err = controllers.GetAllEmails(intPage, intMax)
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("error processing request: %v", err), http.StatusInternalServerError)
		return
	}

	jsonData, err = json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("error marshaling response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}

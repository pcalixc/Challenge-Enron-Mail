package server

import (
	"challenge/api/config"
	"challenge/api/controllers"
	"challenge/api/models"
	"challenge/api/utils"
	"encoding/json"
	"fmt"
	"net/http"

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
	var (
		result   models.HitsResponse
		jsonData []byte
		err      error
		params   models.QueryParameters
	)
	params = utils.EvaluateQueryParams(r.URL.Query().Get("max"), r.URL.Query().Get("page"), r.URL.Query().Get("term"))

	if params.Term != "" {
		result, err = controllers.SearchMails(params.Term, params.PageNumber, params.MaxResults)
	} else {
		result, err = controllers.GetAllEmails(params.PageNumber, params.MaxResults)
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

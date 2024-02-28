package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) setupRoutes() {
	s.ChiRouter.Route("/", func(r chi.Router) {
		//Welcome route
		r.Get("/", s.Welcome)

		//Route in charge of handling emails
		r.Get("/emails", s.handleEmails)
	})
	s.ChiRouter.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
	})
}

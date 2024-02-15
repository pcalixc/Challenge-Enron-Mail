package server

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) setupRoutes() {
	s.ChiRouter.Route("/", func(r chi.Router) {
		r.Get("/", s.Welcome)
		r.Get("/emails", s.handleEmails)

	})
}

package server

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) setupRoutesb() {
	s.ChiRouter.Route("/", func(r chi.Router) {
		r.Get("/", s.Welcome)
		r.Get("/search", s.handleEmailSearch)
		r.Get("/emails", s.handleEmails)
	})
}

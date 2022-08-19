package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"

	"github.com/Timurshk/internal/handlers"
)

type Server struct {
	host     string
	port     string
	handlers *handlers.Handler
	router   *chi.Mux
}

func New(host, port string) *Server {
	return &Server{
		host:     host,
		port:     port,
		handlers: handlers.New(),
		router:   chi.NewRouter(),
	}
}

func (s *Server) Start() {
	handlers := s.handlers
	router := s.router
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Route("/", func(r chi.Router) {
		router.Get("/{id}", handlers.GetURL)
		router.Post("/", handlers.PostURL)
	})
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	log.Fatal(http.ListenAndServe(addr, s.router))
}

package server

import (
	"net/http"
	"sales-register/internal/handler"
)

type ServerCfg struct {
	Port string
}

type Server struct {
	ServerCfg
}

func NewServer(cfg ServerCfg) *Server {
	return &Server{
		ServerCfg: cfg,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("GET /health", handler.Health)
	http.HandleFunc("POST /account", handler.Account)

	return http.ListenAndServe(s.Port, nil)
}

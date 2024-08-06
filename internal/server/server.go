package server

import (
	"context"
	"net/http"
	db "sales-register/db/sqlc"
	"sales-register/internal/handler"
)

type ServerCfg struct {
	Port    string
	Queries *db.Queries
}

type Server struct {
	ServerCfg
}

func NewServer(cfg ServerCfg) *Server {
	return &Server{
		ServerCfg: cfg,
	}
}

func (s *Server) Start(ctx context.Context) error {
	http.HandleFunc("GET /health", handler.Health)
	http.HandleFunc("POST /account", handler.Account(ctx, s.Queries))

	return http.ListenAndServe(s.Port, nil)
}

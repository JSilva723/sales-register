package server

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	return http.ListenAndServe(s.Port, nil)
}

package server

import (
    "net/http"
    "feed-streaming-server/internal/config"
    "feed-streaming-server/internal/stream"
)

type Server struct {
    Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
    return &Server{Config: cfg}
}

func (s *Server) Start() error {
    http.HandleFunc("/stream", stream.HandleStream)
    http.Handle("/", http.FileServer(http.Dir("./web/receiver")))
    return http.ListenAndServe(":"+s.Config.ServerPort, nil)
}

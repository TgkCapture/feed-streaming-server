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
    // Serve static files for sender and receiver
    http.Handle("/sender/", http.StripPrefix("/sender/", http.FileServer(http.Dir("./web/sender"))))
    http.Handle("/receiver/", http.StripPrefix("/receiver/", http.FileServer(http.Dir("./web/receiver"))))
    
    // Handle streaming
    http.HandleFunc("/stream", stream.HandleStream)

    return http.ListenAndServe(":"+s.Config.ServerPort, nil)
}

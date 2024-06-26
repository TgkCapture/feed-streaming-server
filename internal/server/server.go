package server

import (
    "fmt"
    "net/http"
    "feed-streaming-server/internal/config"
    "feed-streaming-server/internal/stream"
    "feed-streaming-server/internal/utils"
)

type Server struct {
    Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
    return &Server{Config: cfg}
}

func (s *Server) Start(role string) error {
    utils.InitLogger()

    // Serve static files for sender and receiver
    http.Handle("/sender/", http.StripPrefix("/sender/", http.FileServer(http.Dir("./web/sender"))))
    http.Handle("/receiver/", http.StripPrefix("/receiver/", http.FileServer(http.Dir("./web/receiver"))))

    // Handle streaming
    http.HandleFunc("/stream", stream.HandleStream)

    var addr string
    if role == "sender" {
        addr = fmt.Sprintf(":%s", s.Config.SenderPort)
        utils.InfoLogger.Printf("Sender server starting on port %s...", s.Config.SenderPort)
    } else if role == "receiver" {
        addr = fmt.Sprintf(":%s", s.Config.ReceiverPort)
        utils.InfoLogger.Printf("Receiver server starting on port %s...", s.Config.ReceiverPort)
    } else {
        return fmt.Errorf("unknown server role: %s", role)
    }

    if err := http.ListenAndServe(addr, nil); err != nil {
        return fmt.Errorf("error starting %s server: %v", role, err)
    }

    return nil
}

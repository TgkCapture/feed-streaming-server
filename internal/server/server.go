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

func (s *Server) Start() error {
    utils.InitLogger()

    // Serve static files for sender and receiver
    http.Handle("/sender/", http.StripPrefix("/sender/", http.FileServer(http.Dir("./web/sender"))))
    http.Handle("/receiver/", http.StripPrefix("/receiver/", http.FileServer(http.Dir("./web/receiver"))))
    
    // Handle streaming
    http.HandleFunc("/stream", stream.HandleStream)

    senderAddr := fmt.Sprintf(":%s", s.Config.SenderPort)
    receiverAddr := fmt.Sprintf(":%s", s.Config.ReceiverPort)

    go func() {
        if err := http.ListenAndServe(senderAddr, nil); err != nil {
            utils.ErrorLogger.Fatalf("Error starting sender server: %v", err)
        }
    }()

    utils.InfoLogger.Printf("Sender server starting on port %s...", s.Config.SenderPort)

    if err := http.ListenAndServe(receiverAddr, nil); err != nil {
        utils.ErrorLogger.Fatalf("Error starting receiver server: %v", err)
    }

    utils.InfoLogger.Printf("Receiver server starting on port %s...", s.Config.ReceiverPort)


    return nil
}

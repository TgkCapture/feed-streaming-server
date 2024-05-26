package main

import (
    "log"
    "feed-streaming-server/internal/config"
    "feed-streaming-server/internal/server"
)

func main() {
    cfg := config.LoadConfig()
    srv := server.NewServer(cfg)

    log.Printf("Sender server starting on port %s...", cfg.ServerPort)
    if err := srv.Start(); err != nil {
        log.Fatalf("Error starting sender server: %v", err)
    }
}

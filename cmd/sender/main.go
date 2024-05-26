package main

import (
    "feed-streaming-server/internal/config"
    "feed-streaming-server/internal/server"
    "feed-streaming-server/internal/utils"
)

func main() {
    utils.InitLogger()
    
    cfg := config.LoadConfig()
    srv := server.NewServer(cfg)

    utils.InfoLogger.Printf("Sender server starting on port %s...", cfg.SenderPort)
    if err := srv.Start(); err != nil {
        utils.ErrorLogger.Fatalf("Error starting sender server: %v", err)
    }
}

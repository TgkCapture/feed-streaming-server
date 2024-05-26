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

    utils.InfoLogger.Printf("Receiver server starting on port %s...", cfg.ReceiverPort)
    if err := srv.Start(); err != nil {
        utils.ErrorLogger.Fatalf("Error starting receiver server: %v", err)
    }
}

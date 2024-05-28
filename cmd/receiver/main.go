package main

import (
    "github.com/TgkCapture/feed-streaming-server/internal/config"
    "github.com/TgkCapture/feed-streaming-server/internal/server"
    "github.com/TgkCapture/feed-streaming-server/internal/utils"
)

func main() {
    utils.InitLogger()

    cfg := config.LoadConfig()
    srv := server.NewServer(cfg)

    utils.InfoLogger.Printf("Receiver server starting on port %s...", cfg.ReceiverPort)
    if err := srv.Start("receiver"); err != nil {
        utils.ErrorLogger.Fatalf("Error starting receiver server: %v", err)
    }
}

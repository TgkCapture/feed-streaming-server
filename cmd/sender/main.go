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

    utils.InfoLogger.Printf("Sender server starting on port %s...", cfg.SenderPort)
    if err := srv.Start("sender"); err != nil {
        utils.ErrorLogger.Fatalf("Error starting sender server: %v", err)
    }
}

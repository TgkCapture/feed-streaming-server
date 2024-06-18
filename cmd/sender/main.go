package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TgkCapture/feed-streaming-server/internal/config"
	"github.com/TgkCapture/feed-streaming-server/internal/db"
	"github.com/TgkCapture/feed-streaming-server/internal/server"
	"github.com/TgkCapture/feed-streaming-server/internal/utils"
)

func main() {
    utils.InitLogger()
    
    cfg := config.LoadConfig()

    if err := db.InitDB(cfg); err != nil {
        utils.ErrorLogger.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.CloseDB()

    srv := server.NewServer(cfg)

    // Handle graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-quit
        log.Println("Shutting down server...")
        db.CloseDB()
        os.Exit(0)
    }()

    utils.InfoLogger.Printf("Sender server starting on port %s...", cfg.SenderPort)
    if err := srv.Start("sender"); err != nil {
        utils.ErrorLogger.Fatalf("Error starting sender server: %v", err)
    }
}

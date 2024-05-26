package config

import (
	"feed-streaming-server/internal/utils"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
    ServerPort string
    StreamKey  string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        utils.ErrorLogger.Fatalf("Error loading .env file")
    }

    return &Config{
        ServerPort: os.Getenv("SERVER_PORT"),
        StreamKey:  os.Getenv("STREAM_KEY"),
    }
}

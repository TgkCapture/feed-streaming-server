package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    ServerPort string
    StreamKey  string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    return &Config{
        ServerPort: os.Getenv("SERVER_PORT"),
        StreamKey:  os.Getenv("STREAM_KEY"),
    }
}

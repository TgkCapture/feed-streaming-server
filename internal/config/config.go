package config

import (
    "github.com/TgkCapture/feed-streaming-server/internal/utils"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    SenderPort   string
    ReceiverPort string
    StreamKey    string
    DBUser       string
    DBPassword   string
    DBName       string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        utils.ErrorLogger.Fatalf("Error loading .env file")
    }

    return &Config{
        SenderPort:   os.Getenv("SENDER_PORT"),
        ReceiverPort: os.Getenv("RECEIVER_PORT"),
        StreamKey:    os.Getenv("STREAM_KEY"),
        DBUser:       os.Getenv("DB_USER"),
        DBPassword:   os.Getenv("DB_PASSWORD"),
        DBName:       os.Getenv("DB_NAME"),
    }
}

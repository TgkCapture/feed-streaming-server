package db

import (
    "database/sql"
    "fmt"
    "github.com/TgkCapture/feed-streaming-server/internal/config"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var DB *sql.DB

func InitDB(cfg *config.Config) error {
    dsn := fmt.Sprintf("%s:%s@/%s", cfg.DBUser, cfg.DBPassword, cfg.DBName)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("error opening database: %w", err)
    }

    if err = db.Ping(); err != nil {
        return fmt.Errorf("error connecting to the database: %w", err)
    }

    DB = db
    return nil
}

func CloseDB() {
    if DB != nil {
        if err := DB.Close(); err != nil {
            log.Printf("error closing database: %v", err)
        }
    }
}

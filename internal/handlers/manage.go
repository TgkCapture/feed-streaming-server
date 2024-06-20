package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/TgkCapture/feed-streaming-server/internal/db"
    "github.com/TgkCapture/feed-streaming-server/internal/utils"
)

type Video struct {
    ID       int    `json:"id"`
    Filename string `json:"filename"`
    URL      string `json:"url"`
}

func GetAllVideosHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, filename, url FROM videos")
    if err != nil {
        utils.ErrorLogger.Printf("Error fetching videos: %v", err)
        http.Error(w, "Error fetching videos", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var videos []Video
    for rows.Next() {
        var video Video
        if err := rows.Scan(&video.ID, &video.Filename, &video.URL); err != nil {
            utils.ErrorLogger.Printf("Error scanning video: %v", err)
            http.Error(w, "Error scanning video", http.StatusInternalServerError)
            return
        }
        videos = append(videos, video)
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(videos); err != nil {
        utils.ErrorLogger.Printf("Error encoding videos: %v", err)
        http.Error(w, "Error encoding videos", http.StatusInternalServerError)
        return
    }
}

func DeleteVideoHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID parameter missing", http.StatusBadRequest)
        return
    }

    _, err := db.DB.Exec("DELETE FROM videos WHERE id = ?", id)
    if err != nil {
        utils.ErrorLogger.Printf("Error deleting video: %v", err)
        http.Error(w, "Error deleting video", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}


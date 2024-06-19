package stream

import (
    "database/sql"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"

    "github.com/TgkCapture/feed-streaming-server/internal/utils"
    "github.com/TgkCapture/feed-streaming-server/internal/db"
)

func HandleStream(w http.ResponseWriter, r *http.Request) {
    if db.DB == nil {
        http.Error(w, "Database not initialized", http.StatusInternalServerError)
        return
    }

    if r.Method == http.MethodPost {
        file, header, err := r.FormFile("video")
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        defer file.Close()

        title := r.FormValue("title")
        description := r.FormValue("description")

        filename := utils.GenerateUniqueFilename(header.Filename)
        uploadDir := "./internal/utils/uploaded_videos"
        if err := ensureDir(uploadDir); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        filePath := filepath.Join(uploadDir, filename)
        f, err := os.Create(filePath)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer f.Close()
        _, err = io.Copy(f, file)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fileURL := fmt.Sprintf("/uploaded_videos/%s", filename)

        // Insert into database
        _, err = db.DB.Exec("INSERT INTO videos (filename, url, title, description) VALUES (?, ?, ?, ?)",
            filename, fileURL, title, description)
        if err != nil {
            http.Error(w, "Database insert error", http.StatusInternalServerError)
            return
        }

        utils.InfoLogger.Printf("Uploaded video saved as: %s", filename)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Uploaded video saved successfully"))

    } else if r.Method == http.MethodGet {
        // Retrieve the latest video from the database
        var filename, url string
        err := db.DB.QueryRow("SELECT filename, url FROM videos ORDER BY upload_time DESC LIMIT 1").Scan(&filename, &url)
        if err != nil {
            if err == sql.ErrNoRows {
                http.Error(w, "No videos found", http.StatusNotFound)
            } else {
                http.Error(w, "Database query error", http.StatusInternalServerError)
            }
            return
        }

        // Stream the video file
        filePath := filepath.Join("./internal/utils", url)
        file, err := os.Open(filePath)
        if err != nil {
            http.Error(w, "File not found", http.StatusNotFound)
            return
        }
        defer file.Close()

        fi, err := file.Stat()
        if err != nil {
            http.Error(w, "File not found", http.StatusNotFound)
            return
        }

        http.ServeContent(w, r, fi.Name(), fi.ModTime(), file)
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func ensureDir(dirName string) error {
    err := os.MkdirAll(dirName, os.ModePerm)
    if err != nil {
        return fmt.Errorf("failed to create directory: %w", err)
    }
    return nil
}

package stream

import (
	"feed-streaming-server/internal/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func HandleStream(w http.ResponseWriter, r *http.Request) {
    uploadDir := "./internal/utils/uploaded_videos"
    if err := ensureDir(uploadDir); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if r.Method == http.MethodPost {
        // Handle incoming stream
        file, header, err := r.FormFile("video")
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        defer file.Close()

        // Generate a unique filename
        filename := utils.GenerateUniqueFilename(header.Filename)

        f, err := os.Create(filepath.Join(uploadDir, filename))
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

        utils.InfoLogger.Printf("Uploaded video saved as: %s", filename)

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Uploaded video saved successfully"))
    } else if r.Method == http.MethodGet {
        // Stream to the viewer the lastest file
        latestFile, err := getLatestFile(uploadDir)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        http.ServeFile(w, r, filepath.Join(uploadDir, latestFile))
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

func getLatestFile(dirPath string) (string, error) {
    var latestModTime time.Time
    var latestFile string

    err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && info.ModTime().After(latestModTime) {
            latestModTime = info.ModTime()
            latestFile = info.Name()
        }
        return nil
    })
    if err != nil {
        return "", err
    }

    return latestFile, nil
}

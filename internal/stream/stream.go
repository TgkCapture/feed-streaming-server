package stream

import (
    "io"
    "log"
    "net/http"
)

func HandleStream(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Handle incoming stream
        file, _, err := r.FormFile("video")
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        defer file.Close()

        // Simulate saving or broadcasting the video stream
        log.Println("Receiving stream")
        io.Copy(io.Discard, file)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Stream received"))
    } else if r.Method == http.MethodGet {
        // Stream to the viewer
        http.ServeFile(w, r, "./internal/utils/nature.mp4")
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

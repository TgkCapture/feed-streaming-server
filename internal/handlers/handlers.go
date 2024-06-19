package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/TgkCapture/feed-streaming-server/internal/utils"
    "github.com/TgkCapture/feed-streaming-server/internal/db"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := utils.VerifyCredentials(db.DB, req.Username, req.Password); err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(req.Username)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    res := LoginResponse{Token: token}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

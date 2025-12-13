package handlers

import (
	"encoding/json"
	"net/http"
	"myapp/internal/auth"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=2"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Username == "" {
		http.Error(w, "username required", http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, "token error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

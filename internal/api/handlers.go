package api

import (
	"context"
	"encoding/json"
	"matchmaking-service/internal/app"
	"matchmaking-service/internal/data"
	"net/http"
	"time"
)

type Handler struct {
	service *app.MatchmakingService
}

func NewHandler(service *app.MatchmakingService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) JoinMatchmaking(w http.ResponseWriter, r *http.Request) {
	var player data.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	h.service.JoinMatchmaking(ctx, player)
	w.WriteHeader(http.StatusAccepted)
}

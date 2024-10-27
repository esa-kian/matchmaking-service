package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"matchmaking-service/internal/api"
	"matchmaking-service/internal/app"
	"matchmaking-service/internal/data"
)

func main() {
	storage := data.NewInMemoryStorage()
	service := app.NewMatchmakingService(storage)
	handler := api.NewHandler(service)

	// Start the background process to check for pending competitions
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			service.StartPendingCompetitions(context.Background(), 30*time.Second)
		}
	}()

	http.HandleFunc("/matchmaking/join", handler.JoinMatchmaking)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

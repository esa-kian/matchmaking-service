package app

import (
	"context"
	"time"

	"matchmaking-service/internal/data"
)

type MatchmakingService struct {
	storage data.Storage
}

func NewMatchmakingService(storage data.Storage) *MatchmakingService {
	return &MatchmakingService{
		storage: storage,
	}
}

func (s *MatchmakingService) JoinMatchmaking(ctx context.Context, player data.Player) {
	s.storage.EnqueuePlayer(player)
	s.CheckForMatchmaking(ctx)
}

func (s *MatchmakingService) CheckForMatchmaking(ctx context.Context) {
	queue := s.storage.GetMatchmakingQueue()

	if len(queue) < 10 {
		return // Not enough players
	}

	// Select players for a new competition
	competition := data.Competition{
		ID:        generateCompetitionID(),
		Players:   queue[:10],
		CreatedAt: time.Now(),
	}
	s.storage.SaveCompetition(competition)
}

func (s *MatchmakingService) StartPendingCompetitions(ctx context.Context, maxWait time.Duration) {
	s.storage.StartPendingCompetitions(maxWait)
}

func generateCompetitionID() string {
	return "unique-id" // This should generate a unique identifier (e.g., UUID)
}

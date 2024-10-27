package data

import (
	"sync"
	"time"
)

type Storage interface {
	EnqueuePlayer(player Player)
	GetMatchmakingQueue() []Player
	SaveCompetition(competition Competition)
	StartPendingCompetitions(maxWaitingTime time.Duration)
}

type InMemoryStorage struct {
	mu           sync.Mutex
	playerQueue  []Player
	competitions []Competition
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		playerQueue:  make([]Player, 0),
		competitions: make([]Competition, 0),
	}
}

func (s *InMemoryStorage) EnqueuePlayer(player Player) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.playerQueue = append(s.playerQueue, player)
}

func (s *InMemoryStorage) GetMatchmakingQueue() []Player {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.playerQueue
}

func (s *InMemoryStorage) SaveCompetition(competition Competition) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.competitions = append(s.competitions, competition)
}

func (s *InMemoryStorage) StartPendingCompetitions(maxWaitingTime time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()

	for i, comp := range s.competitions {
		if !comp.IsStarted && now.Sub(comp.CreatedAt) > maxWaitingTime {
			s.competitions[i].IsStarted = true
			// Notify players (this can be expanded in real implementation)
		}
	}
}

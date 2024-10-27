package data

import "time"

type Player struct {
	ID      string `json:"id"`
	Level   int    `json:"level"`
	Country string `json:"country"`
}

type Competition struct {
	ID        string    `json:"id"`
	Players   []Player  `json:"players"`
	CreatedAt time.Time `json:"created_at"`
	IsStarted bool      `json:"is_started"`
}

package main

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: make(map[string]int)}
}

func (s *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return s.store[player]
}

func (s *InMemoryPlayerStore) RecordWin(player string) {
	s.mu.Lock()
	s.store[player]++
	s.mu.Unlock()
}

func (s *InMemoryPlayerStore) UpdateScore(player string, score int) {
	s.mu.Lock()
	s.store[player] = score
	s.mu.Unlock()
}

func (s *InMemoryPlayerStore) GetLeague() (league []Player) {
	for name, score := range s.store {
		league = append(league, Player{
			Name: name,
			Wins: score,
		})
	}

	return
}

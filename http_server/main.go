package main

import (
	"log"
	"net/http"
	"sync"
)

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

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	log.Println("listing on port :8000")
	log.Fatal(http.ListenAndServe(":8000", server))
}

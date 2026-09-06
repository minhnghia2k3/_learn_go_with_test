package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	UpdateScore(name string, score int)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	server := new(PlayerServer)

	server.store = store

	router := http.NewServeMux()
	router.HandleFunc("/league", server.leagueHandler)
	router.HandleFunc("/players/", server.playerHandler)

	server.Handler = router

	return server
}

func (s *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(s.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

func (s *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/players/"), "/")
	name := parts[0]

	switch method {
	case http.MethodGet:
		s.showScore(w, name)
	case http.MethodPost:
		s.processWin(w, name)
	case http.MethodPut:
		if len(parts) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		score, _ := strconv.Atoi(parts[1])
		s.updateScore(w, name, score)
	}
}

func (s *PlayerServer) showScore(w http.ResponseWriter, name string) {
	score := s.store.GetPlayerScore(name)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (s *PlayerServer) processWin(w http.ResponseWriter, name string) {
	s.store.RecordWin(name)
	w.WriteHeader(http.StatusCreated)
}

func (s *PlayerServer) updateScore(w http.ResponseWriter, name string, score int) {
	s.store.UpdateScore(name, score)
	w.WriteHeader(http.StatusOK)
}

package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	name := strings.TrimPrefix(r.URL.Path, "/players/")

	switch method {
	case http.MethodGet:
		s.showScore(w, name)
	case http.MethodPost:
		s.processWin(w, name)
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

func (s *PlayerServer) GetPlayerScore(name string) int {
	if name == "Floyd" {
		return 10
	}

	if name == "Pepper" {
		return 20
	}

	return 0
}

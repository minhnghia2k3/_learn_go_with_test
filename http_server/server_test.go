package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayers(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{store}

	t.Run("Return Pepper's score", func(t *testing.T) {
		req := newGetScoreRequest(t, "Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusOK)
		assertScore(t, res.Body.String(), "20")
	})

	t.Run("Return Floyd's score", func(t *testing.T) {
		req := newGetScoreRequest(t, "Floyd")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusOK)
		assertScore(t, res.Body.String(), "10")
	})

	t.Run("Return 404 if not found player", func(t *testing.T) {
		req := newGetScoreRequest(t, "Apollo")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreScores(t *testing.T) {
	store := &StubPlayerStore{
		scores:   make(map[string]int),
		winCalls: nil,
	}
	server := &PlayerServer{store}

	t.Run("should returns created on POST", func(t *testing.T) {
		name := "Pepper"

		req := newPostWinRequest(t, "Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusCreated)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls, expected %d calls", len(store.winCalls), 1)
		}

		if store.winCalls[0] != name {
			t.Errorf("got name %q, expected name %q", store.winCalls[0], name)
		}
	})
}

func newGetScoreRequest(t *testing.T, player string) *http.Request {
	t.Helper()

	path := fmt.Sprintf("/players/%s", player)

	request, _ := http.NewRequest(http.MethodGet, path, nil)
	return request
}

func newPostWinRequest(t *testing.T, player string) *http.Request {
	t.Helper()

	path := fmt.Sprintf("/players/%s", player)

	request, _ := http.NewRequest(http.MethodPost, path, nil)

	return request
}

func assertScore(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

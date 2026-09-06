package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

const jsonContentType = "application/json"

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) UpdateScore(name string, score int) {
	s.scores[name] = score
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGetPlayers(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := NewPlayerServer(store)

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
	server := NewPlayerServer(store)

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

func TestUpdateScores(t *testing.T) {
	store := &StubPlayerStore{
		scores: make(map[string]int),
	}
	srv := NewPlayerServer(store)
	t.Run("should update player score", func(t *testing.T) {
		player := "Pepper"
		want := "3"
		path := fmt.Sprintf("/players/%s/%s", player, want)

		request, _ := http.NewRequest(http.MethodPut, path, nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)

		playerScore := srv.store.GetPlayerScore(player)
		// convert int to string
		got := strconv.Itoa(playerScore)

		assertScore(t, got, want)
	})

	t.Run("should not update when no score", func(t *testing.T) {
		player := "Pepper"
		path := fmt.Sprintf("/players/%s", player)

		request, _ := http.NewRequest(http.MethodPut, path, nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusBadRequest)
	})
}

func TestLeague(t *testing.T) {

	t.Run("it return 200 on /league", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := &StubPlayerStore{scores: nil, winCalls: nil, league: wantedLeague}
		server := NewPlayerServer(store)

		request := newGetLeague(t)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, jsonContentType)
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

func newGetLeague(t *testing.T) *http.Request {
	t.Helper()

	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Errorf("Unable to parse response from server %q: %v", body, err)
	}
	return
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

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got league %v, want league %v", got, want)
	}
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	contentType := response.Result().Header.Get("content-type")
	if contentType != want {
		t.Errorf("expected want type = %q, got %v", want, contentType)
	}
}

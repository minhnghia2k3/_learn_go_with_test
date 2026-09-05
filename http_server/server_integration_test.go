package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecodingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}

	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(t, player))

	response := httptest.NewRecorder()
	request := newGetScoreRequest(t, player)
	server.ServeHTTP(response, request)

	assertStatusCode(t, response.Code, http.StatusOK)
	assertScore(t, response.Body.String(), "3")
}

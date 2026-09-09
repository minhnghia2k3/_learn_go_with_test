package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecodingWinsAndRetrievingThem(t *testing.T) {
	database, removeFile := createTempFile(t, `[]`)
	defer removeFile()
	store, err := NewFileSystemPLayerStore(database)
	if err != nil {
		t.Fatalf("did'nt expect an error but got one, %v", err)
	}
	server := NewPlayerServer(store)

	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(t, player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetScoreRequest(t, player)
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertScore(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetLeague(t)
		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{Name: "Pepper", Wins: 3},
		}

		assertLeague(t, got, want)
	})
}

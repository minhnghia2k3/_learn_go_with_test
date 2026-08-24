package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speed of servers", func(t *testing.T) {
		slowServer := makeServer(20 * time.Millisecond)
		fastServer := makeServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := Racer(slowServer.URL, fastServer.URL)

		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}

	})

	t.Run("returns error if request takes more than 10s", func(t *testing.T) {
		slowUrl := makeServer(11 * time.Millisecond)
		fastUrl := makeServer(12 * time.Millisecond)

		_, err := ConfigurableRacer(slowUrl.URL, fastUrl.URL, 10*time.Millisecond)

		if err == nil {
			t.Fatal("expected an error")
		}
	})
}

func makeServer(sleep time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleep)
		w.WriteHeader(http.StatusOK)
	}))

	return server
}

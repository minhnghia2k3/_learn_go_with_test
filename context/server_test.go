package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStoreWriter struct {
	written bool
}

func (w *SpyStoreWriter) Header() http.Header {
	w.written = true
	return nil
}
func (w *SpyStoreWriter) Write([]byte) (int, error) {
	w.written = true
	return 0, errors.New("")
}
func (w *SpyStoreWriter) WriteHeader(statusCode int) {
	w.written = true
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string)

	go func() {
		// loop through response body's character
		var result string

		select {
		case <-ctx.Done():
			log.Println("context is cancelled")
			return
		default:
			for _, c := range s.response {
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		// send data out to data chan
		data <- result
	}()

	// race between ctx.Done() and data channel
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case d := <-data:
		return d, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q, want %q", response.Body.String(), data)
		}
	})
	t.Run("tell store to stop the work, if request is cancelled", func(t *testing.T) {
		data := "Hello world!"
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// incoming request create a ctx
		cancelCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancelCtx)
		response := &SpyStoreWriter{}

		// outgoing accept a ctx
		srv.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response should not be written")
		}
	})
}

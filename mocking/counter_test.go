package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

const (
	sleep = "sleep"
	write = "write"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyOperations struct {
	Calls []string
}

func (s *SpyOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyOperations) Write(p []byte) (int, error) {
	s.Calls = append(s.Calls, write)
	return 0, nil
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountDown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Counter(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if sleeper.Calls != 3 {
			t.Errorf("not enough sleeper called, got %d want 3", sleeper.Calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spyOperations := SpyOperations{}

		Counter(&spyOperations, &spyOperations)

		got := spyOperations.Calls
		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !slices.Equal(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func TestConfigurableSleep(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}

	sleeper.Sleep()

	if sleepTime != spyTime.durationSlept {
		t.Errorf("should slept %v, but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

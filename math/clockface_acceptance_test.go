package clockface_test

import (
	"testing"
	"time"

	clockface "github.com/minhnghia2k3/_learn_go_with_test/math"
)

func TestSecondHandMidnight(t *testing.T) {
	tm := time.Date(1337, 1, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, 1, 1, 0, 0, 30, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(tm)

	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}

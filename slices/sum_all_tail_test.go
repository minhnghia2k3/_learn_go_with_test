package main

import (
	"slices"
	"testing"
)

func TestSumAllTail(t *testing.T) {
	checkSlicesEquals := func(t testing.TB, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got '%d', want '%d'", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{3, 5})
		want := []int{2, 5}

		checkSlicesEquals(t, got, want)
	})

	t.Run("make the sums of empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSlicesEquals(t, got, want)
	})

}

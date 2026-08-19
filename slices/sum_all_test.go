package main

import (
	"slices"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	if !slices.Equal(got, want) {
		t.Errorf("got '%d', want '%d'", got, want)
	}
}

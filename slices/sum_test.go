package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got '%d', want '%d', %v", got, want, numbers)
		}
	})

}

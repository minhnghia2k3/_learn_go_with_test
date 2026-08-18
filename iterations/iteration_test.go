package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat the character 6 times", func(t *testing.T) {
		got := Repeat("a", 6)
		want := "aaaaaa"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func BenchmarkRepeat(t *testing.B) {
	for t.Loop() {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 5)
	fmt.Println(res)
	// Output: aaaaa
}

package main

import (
	"maps"
	"testing"
	"time"
)

func mockWebChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestWebChecker(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsite(mockWebChecker, websites)

	if !maps.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}

func slowMockWebChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebChecker(b *testing.B) {
	urls := make([]string, 100)

	for i := range 100 {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsite(slowMockWebChecker, urls)
	}
}

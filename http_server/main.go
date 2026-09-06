package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	log.Println("listing on port :8000")
	log.Fatal(http.ListenAndServe(":8000", server))
}

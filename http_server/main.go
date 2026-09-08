package main

import (
	"log"
	"net/http"
	"os"
)

const gameDatabase = "game.db.json"

func main() {
	file, err := os.OpenFile(gameDatabase, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("openning file got error:", err)
	}

	store, err := NewFileSystemPLayerStore(file)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := NewPlayerServer(store)

	log.Println("listing on port :8000")
	log.Fatal(http.ListenAndServe(":8000", server))
}
